package repository

import (
	"context"
	"database/sql"
	"log/slog"
	"webook-server/internal/domain"
	"webook-server/internal/model"
	"webook-server/internal/repository/cache"
	"webook-server/internal/repository/dao"
)

type UserRepository interface {
	Create(ctx context.Context, u domain.User) error
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindByUserId(ctx context.Context, userId int64) (domain.User, error)
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

func (r *UserRepositoryByGormAndRedis) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, toModel(u))
}

func (r *UserRepositoryByGormAndRedis) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	u, err := r.dao.FindByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, err
	}
	return toDomain(u), nil
}

func (r *UserRepositoryByGormAndRedis) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return toDomain(u), nil
}

func (r *UserRepositoryByGormAndRedis) FindByUserId(ctx context.Context, userId int64) (domain.User, error) {
	//查询缓存
	if u, err := r.cache.Get(ctx, userId); err == nil {
		return u, nil
	}
	//缓存不命中，查询数据库
	u, err := r.dao.FindByUserId(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}
	domainUser := toDomain(u)

	go func() {
		if err = r.cache.Set(ctx, domainUser); err != nil { // 直接忽略错误也ok
			//todo log
			slog.Warn("cache domainUser fail", err.Error())
		}
	}()

	return domainUser, nil
}

func toDomain(u model.User) domain.User {
	return domain.User{
		UserId:   u.UserId,
		Email:    u.Email.String,
		Password: u.Password,
		UserName: u.UserName.String,
		Phone:    u.Phone.String,
	}
}

func toModel(u domain.User) model.User {
	return model.User{
		UserId:   u.UserId,
		UserName: sql.NullString{String: u.UserName, Valid: u.UserName != ""},
		Email:    sql.NullString{String: u.Email, Valid: u.Email != ""},
		Phone:    sql.NullString{String: u.Phone, Valid: u.Phone != ""},
		Password: u.Password,
	}
}
