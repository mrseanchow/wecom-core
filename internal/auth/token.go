package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/mrseanchow/wecom-core/pkg/cache"
	"github.com/mrseanchow/wecom-core/pkg/logger"
)

const (
	// TokenExpireOffset token 提前刷新时间（秒�?
	// 提前 5 分钟刷新 token，避免在使用时过�?
	TokenExpireOffset = 300
)

// TokenResponse token响应
type TokenResponse struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// TokenManager Token管理�?
type TokenManager struct {
	// corpID 企业ID
	corpID string
	// agents 多应用配置，key为应用名称或ID
	agents map[string]*AgentInfo
	// baseURL API基础URL
	baseURL string
	// cache 缓存
	cache cache.Cache
	// logger 日志记录�?
	logger logger.Logger
	// httpClient HTTP客户�?
	httpClient *http.Client
	// refreshLock 刷新�?使用map存储每个应用的锁)
	refreshLocks map[string]*sync.Mutex
	// refreshLocksMapLock 保护refreshLocks map的锁
	refreshLocksMapLock sync.Mutex
}

// AgentInfo 应用信息
type AgentInfo struct {
	corpID  string
	AgentID int64
	Secret  string
	Name    string
}

// NewTokenManager 创建Token管理�?
func NewTokenManager(corpID, baseURL string, c cache.Cache, log logger.Logger) *TokenManager {
	if c == nil {
		c = NewMemoryCache()
	}

	return &TokenManager{
		corpID:              corpID,
		baseURL:             baseURL,
		cache:               c,
		logger:              log,
		httpClient:          &http.Client{Timeout: 30 * time.Second},
		agents:              make(map[string]*AgentInfo),
		refreshLocks:        make(map[string]*sync.Mutex),
		refreshLocksMapLock: sync.Mutex{},
	}
}

// RegisterAgent 注册应用
func (tm *TokenManager) RegisterAgent(agentKey string, corpID string, agentID int64, secret string) {
	tm.agents[agentKey] = &AgentInfo{
		corpID:  corpID,
		AgentID: agentID,
		Secret:  secret,
		Name:    agentKey,
	}
}

// getRefreshLock 获取应用的刷新锁
func (tm *TokenManager) getRefreshLock(agentKey string) *sync.Mutex {
	tm.refreshLocksMapLock.Lock()
	defer tm.refreshLocksMapLock.Unlock()

	if lock, exists := tm.refreshLocks[agentKey]; exists {
		return lock
	}

	lock := &sync.Mutex{}
	tm.refreshLocks[agentKey] = lock
	return lock
}

// GetToken 获取token（默认应用，向后兼容�?
func (tm *TokenManager) GetToken(ctx context.Context) (string, error) {
	return tm.GetTokenByAgent(ctx, "")
}

// GetTokenByAgent 根据应用key获取token
func (tm *TokenManager) GetTokenByAgent(ctx context.Context, agentKey string) (string, error) {
	// 获取应用�?secret
	secret := tm.getAgentSecret(agentKey)
	if secret == "" {
		return "", fmt.Errorf("agent not found or secret is empty: %s", agentKey)
	}

	cacheKey := tm.cacheKey(agentKey)

	// 1. 从缓存获�?
	token, expireAt, err := tm.cache.Get(ctx, cacheKey)
	if err == nil && time.Now().Before(expireAt) {
		tm.logger.Debug("Token retrieved from cache",
			logger.F("agent_key", agentKey),
			logger.F("expire_at", expireAt))
		return token, nil
	}

	// 2. 加锁刷新（防止并发重复获取）
	lock := tm.getRefreshLock(agentKey)
	lock.Lock()
	defer lock.Unlock()

	// 3. Double-check（可能已被其他协程刷新）
	token, expireAt, err = tm.cache.Get(ctx, cacheKey)
	if err == nil && time.Now().Before(expireAt) {
		tm.logger.Debug("Token retrieved from cache after lock",
			logger.F("agent_key", agentKey),
			logger.F("expire_at", expireAt))
		return token, nil
	}

	// 4. 调用 API 获取 token
	tm.logger.Info("Fetching new token from API",
		logger.F("agent_key", agentKey))
	token, expiresIn, err := tm.fetchTokenFromAPI(ctx, secret)
	if err != nil {
		tm.logger.Error("Failed to fetch token",
			logger.F("agent_key", agentKey),
			logger.F("error", err))
		return "", err
	}

	// 5. 缓存 token（提�?5 分钟过期�?
	expireAt = time.Now().Add(time.Duration(expiresIn-TokenExpireOffset) * time.Second)
	if err := tm.cache.Set(ctx, cacheKey, token, expireAt); err != nil {
		tm.logger.Warn("Failed to cache token",
			logger.F("agent_key", agentKey),
			logger.F("error", err))
	}

	tm.logger.Info("Token refreshed successfully",
		logger.F("agent_key", agentKey),
		logger.F("expires_in", expiresIn),
		logger.F("expire_at", expireAt))

	return token, nil
}

// RefreshToken 强制刷新token（用�?token 失效重试，默认应用）
func (tm *TokenManager) RefreshToken(ctx context.Context) error {
	return tm.RefreshTokenByAgent(ctx, "")
}

// RefreshTokenByAgent 强制刷新指定应用的token
func (tm *TokenManager) RefreshTokenByAgent(ctx context.Context, agentKey string) error {
	// 获取应用�?secret
	secret := tm.getAgentSecret(agentKey)
	if secret == "" {
		return fmt.Errorf("agent not found or secret is empty: %s", agentKey)
	}

	lock := tm.getRefreshLock(agentKey)
	lock.Lock()
	defer lock.Unlock()

	tm.logger.Info("Force refreshing token",
		logger.F("agent_key", agentKey))

	token, expiresIn, err := tm.fetchTokenFromAPI(ctx, secret)
	if err != nil {
		tm.logger.Error("Failed to refresh token",
			logger.F("agent_key", agentKey),
			logger.F("error", err))
		return err
	}

	cacheKey := tm.cacheKey(agentKey)
	expireAt := time.Now().Add(time.Duration(expiresIn-TokenExpireOffset) * time.Second)
	if err := tm.cache.Set(ctx, cacheKey, token, expireAt); err != nil {
		tm.logger.Warn("Failed to cache refreshed token",
			logger.F("agent_key", agentKey),
			logger.F("error", err))
	}

	tm.logger.Info("Token force refreshed successfully",
		logger.F("agent_key", agentKey),
		logger.F("expires_in", expiresIn),
		logger.F("expire_at", expireAt))

	return nil
}

// getAgentSecret 获取应用�?secret
func (tm *TokenManager) getAgentSecret(agentKey string) string {
	// 如果 agentKey 为空，使用默认应�?
	if agentKey == "" {
		// 如果只有一个应用，使用该应�?
		if len(tm.agents) == 1 {
			for _, agent := range tm.agents {
				return agent.Secret
			}
		}
		return ""
	}

	// 查找指定应用
	if agent, ok := tm.agents[agentKey]; ok {
		return agent.Secret
	}

	return ""
}

// fetchTokenFromAPI 从API获取token
func (tm *TokenManager) fetchTokenFromAPI(ctx context.Context, secret string) (token string, expiresIn int, err error) {
	url := fmt.Sprintf("%s/cgi-bin/gettoken?corpid=%s&corpsecret=%s",
		tm.baseURL, tm.corpID, secret)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := tm.httpClient.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("failed to fetch token: %w", err)
	}
	defer resp.Body.Close()

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", 0, fmt.Errorf("failed to decode response: %w", err)
	}

	if tokenResp.ErrCode != 0 {
		return "", 0, fmt.Errorf("failed to get token: errcode=%d, errmsg=%s",
			tokenResp.ErrCode, tokenResp.ErrMsg)
	}

	if tokenResp.AccessToken == "" {
		return "", 0, errors.New("empty access_token in response")
	}

	return tokenResp.AccessToken, tokenResp.ExpiresIn, nil
}

// cacheKey 获取缓存key
func (tm *TokenManager) cacheKey(agentKey string) string {
	if agentKey == "" {
		return fmt.Sprintf("wecom:token:%s", tm.corpID)
	}
	return fmt.Sprintf("wecom:token:%s:%s", tm.corpID, agentKey)
}

