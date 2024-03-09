package repository

import (
	"context"
	"webook-server/internal/domain"
	"webook-server/internal/repository/dao"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrRecordNotFound
)

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{dao: dao}
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
	u, err := r.dao.QueryByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return toDomain(u), nil
}

func (r *UserRepository) FindByUserId(ctx context.Context, userId int64) (domain.User, error) {
	u, err := r.dao.QueryByUserId(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}
	return toDomain(u), nil
}

func toDomain(u dao.User) domain.User {
	return domain.User{
		UserId:   u.UserId,
		Email:    u.Email,
		Password: u.Password,
		UserName: u.UserName,
	}
}
