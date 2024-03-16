package repository

import (
	"context"
	"webook-server/internal/repository/cache"
)

var (
	ErrCodeSendFrequently   = cache.ErrCodeSendFrequently
	ErrCodeVerifyFrequently = cache.ErrCodeVerifyFrequently
)

type CodeRepository interface {
	Store(ctx context.Context, biz, phone, code string) error
	Verify(ctx context.Context, biz, phone, code string) (bool, error)
}

type CodeRepositoryByCache struct {
	cache cache.CodeCache
}

func NewCodeRepository(cache cache.CodeCache) CodeRepository {
	return &CodeRepositoryByCache{cache: cache}
}

func (r *CodeRepositoryByCache) Store(ctx context.Context, biz, phone, code string) error {
	return r.cache.Set(ctx, biz, phone, code)
}

func (r *CodeRepositoryByCache) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	return r.cache.Verify(ctx, biz, phone, code)
}
