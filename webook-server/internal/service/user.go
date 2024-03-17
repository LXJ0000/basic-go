package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"webook-server/internal/domain"
	"webook-server/internal/repository"
	"webook-server/pkg/snowflake"
)

var (
	ErrDuplicate             = repository.ErrDuplicate
	ErrInvalidUserOrPassword = errors.New("用户名或密码不正确")
)

type UserService interface {
	Register(ctx context.Context, u domain.User) error
	Login(ctx context.Context, email, password string) (domain.User, error)
	Info(ctx context.Context, userId int64) (domain.User, error)
	FindOrCreate(ctx context.Context, phone string) (domain.User, error)
}

type UserServiceByRepo struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceByRepo{repo: repo}
}

func (svc *UserServiceByRepo) Register(ctx context.Context, u domain.User) error {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(encrypted)
	return svc.repo.Create(ctx, u)
}

func (svc *UserServiceByRepo) Login(ctx context.Context, email, password string) (domain.User, error) {
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (svc *UserServiceByRepo) Info(ctx context.Context, userId int64) (domain.User, error) {
	return svc.repo.FindByUserId(ctx, userId)
}

func (svc *UserServiceByRepo) FindOrCreate(ctx context.Context, phone string) (domain.User, error) {
	//快路径 触发降级操作，只走快路径 即系统资源不不足，只服务已经注册过的用户
	u, err := svc.repo.FindByPhone(ctx, phone)
	if err == nil {
		return u, nil
	}
	u = domain.User{
		UserId: snowflake.GenID(),
		Phone:  phone,
	}
	if err := svc.repo.Create(ctx, u); err != nil {
		return domain.User{}, err
	}
	return u, nil
}
