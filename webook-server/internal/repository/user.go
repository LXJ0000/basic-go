package repository

import (
	"context"
	"log/slog"
	"webook-server/internal/model"
	"webook-server/internal/repository/cache"
	"webook-server/internal/repository/dao"
)

type UserRepository interface {
	Create(ctx context.Context, u model.User) error
	FindByPhone(ctx context.Context, phone string) (model.User, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
	FindByUserId(ctx context.Context, userId int64) (model.User, error)
}

type UserRepositoryByGormAndRedis struct {
	dao   dao.UserDao
	cache cache.UserCache
}

func NewUserRepository(dao dao.UserDao, cache cache.UserCache) UserRepository {
	return &UserRepositoryByGormAndRedis{
		dao:   dao,
		cache: cache,
	}
}

func (r *UserRepositoryByGormAndRedis) Create(ctx context.Context, u model.User) error {
	return r.dao.Insert(ctx, u)
}

func (r *UserRepositoryByGormAndRedis) FindByPhone(ctx context.Context, phone string) (model.User, error) {
	u, err := r.dao.FindByPhone(ctx, phone)
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (r *UserRepositoryByGormAndRedis) FindByEmail(ctx context.Context, email string) (model.User, error) {
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (r *UserRepositoryByGormAndRedis) FindByUserId(ctx context.Context, userId int64) (model.User, error) {
	//查询缓存
	if u, err := r.cache.Get(ctx, userId); err == nil {
		return u, nil
	}
	//缓存不命中，查询数据库
	u, err := r.dao.FindByUserId(ctx, userId)
	if err != nil {
		return model.User{}, err
	}

	go func() {
		if err = r.cache.Set(ctx, u); err != nil { // 直接忽略错误也ok
			//todo log
			slog.Warn("cache modelUser fail", err.Error())
		}
	}()

	return u, nil
}
