package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/shuaidd/wecom-core/pkg/cache"
)

// FileCache 文件缓存实现
type FileCache struct {
	cacheDir string
}

// FileCacheItem 文件缓存项
type FileCacheItem struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}

// NewMemoryCache 创建文件缓存实例
func NewMemoryCache() cache.Cache {
	// 使用当前工作目录的 cache/token 目录
	cacheDir := filepath.Join(".", "cache", "token")
	fmt.Printf("NewMemoryCache: cacheDir = %s\n", cacheDir)

	// 确保缓存目录存在
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		fmt.Printf("NewMemoryCache: os.MkdirAll error = %v\n", err)
	}

	return &FileCache{
		cacheDir: cacheDir,
	}
}

// Get 获取缓存的token
func (fc *FileCache) Get(ctx context.Context, key string) (token string, expireAt time.Time, err error) {
	// 替换key中的冒号为下划线，避免Windows文件系统错误
	key = strings.ReplaceAll(key, ":", "_")
	// 构建缓存文件路径
	cacheFile := filepath.Join(fc.cacheDir, key+".json")

	// 检查文件是否存在
	if _, err := os.Stat(cacheFile); os.IsNotExist(err) {
		return "", time.Time{}, cache.ErrCacheNotFound
	}

	// 读取文件内容
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("read cache file error: %w", err)
	}

	// 解析缓存数据
	var item FileCacheItem
	if err := json.Unmarshal(data, &item); err != nil {
		return "", time.Time{}, fmt.Errorf("unmarshal cache data error: %w", err)
	}

	// 检查是否过期
	if time.Now().After(item.ExpireAt) {
		return "", time.Time{}, cache.ErrCacheExpired
	}

	return item.Token, item.ExpireAt, nil
}

// Set 设置缓存的token
func (fc *FileCache) Set(ctx context.Context, key string, token string, expireAt time.Time) error {
	// 打印缓存目录
	fmt.Printf("FileCache.Set: cacheDir = %s\n", fc.cacheDir)

	// 确保缓存目录存在
	if err := os.MkdirAll(fc.cacheDir, 0755); err != nil {
		fmt.Printf("FileCache.Set: os.MkdirAll error = %v\n", err)
		return fmt.Errorf("create cache directory error: %w", err)
	}

	// 替换key中的冒号为下划线，避免Windows文件系统错误
	key = strings.ReplaceAll(key, ":", "_")
	// 构建缓存文件路径
	cacheFile := filepath.Join(fc.cacheDir, key+".json")
	fmt.Printf("FileCache.Set: cacheFile = %s\n", cacheFile)

	// 创建缓存项
	item := FileCacheItem{
		Token:    token,
		ExpireAt: expireAt,
	}

	// 序列化缓存数据
	data, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		fmt.Printf("FileCache.Set: json.MarshalIndent error = %v\n", err)
		return fmt.Errorf("marshal cache data error: %w", err)
	}

	// 写入文件
	if err := os.WriteFile(cacheFile, data, 0644); err != nil {
		fmt.Printf("FileCache.Set: os.WriteFile error = %v\n", err)
		return fmt.Errorf("write cache file error: %w", err)
	}

	fmt.Printf("FileCache.Set: cache written successfully\n")
	return nil
}

// Delete 删除缓存的token
func (fc *FileCache) Delete(ctx context.Context, key string) error {
	// 替换key中的冒号为下划线，避免Windows文件系统错误
	key = strings.ReplaceAll(key, ":", "_")
	// 构建缓存文件路径
	cacheFile := filepath.Join(fc.cacheDir, key+".json")

	// 检查文件是否存在
	if _, err := os.Stat(cacheFile); os.IsNotExist(err) {
		return nil // 文件不存在，视为删除成功
	}

	// 删除文件
	if err := os.Remove(cacheFile); err != nil {
		return fmt.Errorf("remove cache file error: %w", err)
	}

	return nil
}
