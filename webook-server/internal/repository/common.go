package repository

import (
	"context"
	"webook-server/internal/repository/cache"
	"webook-server/internal/repository/dao"
)

type Repository interface {
	Create(ctx context.Context, key string, val any) error
	Find(ctx context.Context, key string) any
}

type CommonRepository struct {
	dao   dao.Dao
	cache cache.Cache
}

func NewCommonRepository(dao dao.Dao, cache cache.Cache) Repository {
	return &CommonRepository{dao: dao, cache: cache}
}

func (r *CommonRepository) Create(ctx context.Context, key string, val any) error {
	return r.cache.Set(ctx, key, val)
}
func (r *CommonRepository) Find(ctx context.Context, key string) any {
	return r.cache.Get(ctx, key)
}
