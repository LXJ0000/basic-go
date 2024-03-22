package ratelimit

import "context"

type RateLimit interface {
	Limit(ctx context.Context, key string) (bool, error)
}
