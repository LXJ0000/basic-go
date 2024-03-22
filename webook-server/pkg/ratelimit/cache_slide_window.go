package ratelimit

import (
	"context"
	_ "embed"
	"github.com/redis/go-redis/v9"
	"time"
)

//go:embed cache_slice_window.lua
var luaScript string

type CacheSliceWindowLimit struct {
	cmd    redis.Cmdable
	window time.Duration
	rate   int
}

func NewCacheSliceWindowLimiter(cmd redis.Cmdable, window time.Duration, rate int) RateLimit {
	return &CacheSliceWindowLimit{
		cmd:    cmd,
		window: window, // todo
		rate:   rate,   // todo
	}
}

func (l *CacheSliceWindowLimit) Limit(ctx context.Context, key string) (bool, error) {
	return l.cmd.Eval(
		ctx,
		luaScript,
		[]string{key},
		l.window.Milliseconds(),
		l.rate,
		time.Now().UnixMilli(),
	).Bool()
}
