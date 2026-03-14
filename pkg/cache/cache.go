package cache

import (
	"context"
	"errors"
	"time"
)

var (
	// ErrCacheNotFound 缓存未找�?
	ErrCacheNotFound = errors.New("cache not found")

	// ErrCacheExpired 缓存已过�?
	ErrCacheExpired = errors.New("cache expired")
)

// Cache Token缓存接口
type Cache interface {
	// Get 获取缓存的token
	// 返回 token, 过期时间, error
	Get(ctx context.Context, key string) (token string, expireAt time.Time, err error)

	// Set 设置缓存的token
	Set(ctx context.Context, key string, token string, expireAt time.Time) error

	// Delete 删除缓存的token
	Delete(ctx context.Context, key string) error
}

