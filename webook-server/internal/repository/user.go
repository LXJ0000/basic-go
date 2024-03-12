package repository

import (
	"context"
	"log"
	"webook-server/internal/domain"
	"webook-server/internal/repository/cache"
	"webook-server/internal/repository/dao"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrRecordNotFound
)

type UserRepository struct {
	dao   *dao.UserDao
	cache *cache.UserCache
}

func NewUserRepository(dao *dao.UserDao, cache *cache.UserCache) *UserRepository {
	return &UserRepository{
		dao:   dao,
		cache: cache,
	}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		UserId:   u.UserId,
		UserName: u.UserName,
		Email:    u.Email,
		Password: u.Password,
	})
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return toDomain(u), nil
}

func (r *UserRepository) FindByUserId(ctx context.Context, userId int64) (domain.User, error) {
	//查询缓存
	if u, err := r.cache.Get(ctx, userId); err == nil {
		return u, err
	}
	//缓存不命中，查询数据库
	u, err := r.dao.FindByUserId(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}
	domainUser := toDomain(u)

	go func() {
		if err = r.cache.Set(ctx, domainUser); err != nil {
			//todo log
			log.Println("cache domainUser fail")
		}
	}()
	
	return domainUser, nil
}

func toDomain(u dao.User) domain.User {
	return domain.User{
		UserId:   u.UserId,
		Email:    u.Email,
		Password: u.Password,
		UserName: u.UserName,
	}
}
