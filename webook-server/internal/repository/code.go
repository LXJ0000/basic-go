package repository

import (
	"context"
	"webook-server/internal/repository/cache"
)

var (
	ErrCodeSendFrequently   = cache.ErrCodeSendFrequently
	ErrCodeVerifyFrequently = cache.ErrCodeVerifyFrequently
)

type CodeRepository struct {
	cache *cache.CodeCache
}

func NewCodeRepository(cache *cache.CodeCache) *CodeRepository {
	return &CodeRepository{cache: cache}
}

func (r *CodeRepository) Store(ctx context.Context, biz, phone, code string) error {
	return r.cache.Set(ctx, biz, phone, code)
}

func (r *CodeRepository) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	return r.cache.Verify(ctx, biz, phone, code)
}
