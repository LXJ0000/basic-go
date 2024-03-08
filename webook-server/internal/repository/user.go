package repository

import (
	"context"
	"webook-server/internal/domain"
	"webook-server/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{dao: dao}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		UserId:   0, // todo
		UserName: u.UserName,
		Email:    u.Email,
		Password: u.Password,
	})
}
