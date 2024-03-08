package service

import (
	"context"
	"webook-server/internal/domain"
	"webook-server/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (svc *UserService) Register(ctx context.Context, u domain.User) error {
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Login(ctx context.Context) error {
	return nil
}

func (svc *UserService) Profile(ctx context.Context) error {
	return nil
}
