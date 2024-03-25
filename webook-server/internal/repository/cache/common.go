package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set(ctx context.Context, key string, val interface{}) error
	Get(ctx context.Context, key string) any
}

type CommonCache struct {
	cmd redis.Cmdable
}

func NewCommonCache(cmd redis.Cmdable) Cache {
	return &CommonCache{cmd: cmd}
}

func (c *CommonCache) Set(ctx context.Context, key string, val interface{}) error {
	return c.cmd.Set(ctx, key, val, 0).Err() // 不设置过期时间
}

func (c *CommonCache) Get(ctx context.Context, key string) any {
	return c.cmd.Get(ctx, key).String()
}
