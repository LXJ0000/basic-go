package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"testing"
	"webook-server/internal/domain"
	"webook-server/internal/repository"
	repomock "webook-server/internal/repository/mocks"
)

func TestUserServiceByRepo_Login(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		password string

		mock func(ctrl *gomock.Controller) repository.UserRepository

		gotErr error
	}{
		{
			name:     "success",
			email:    "test@qq.com",
			password: "hello@world123",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomock.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "test@qq.com").Return(domain.User{
					Password: "$2a$10$b/4y9QcJwTfcrbTrOWFz8uP0YTiE8e3EJ3jmCh20pl3lVyWnziWEW",
				}, nil)
				return repo
			},
			gotErr: nil,
		},
		{
			name:     "邮箱不存在",
			email:    "err@qq.com",
			password: "hello@world123",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomock.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "err@qq.com").Return(domain.User{}, gorm.ErrRecordNotFound)
				return repo
			},
			gotErr: errors.New("用户名或密码不正确"),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repo := tt.mock(c)
			svc := NewUserService(repo)
			_, err := svc.Login(context.Background(), tt.email, tt.password)
			assert.Equal(t, tt.gotErr, err)
		})
	}
}

func TestGenerateFromPassword(t *testing.T) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte("hello@world123"), bcrypt.DefaultCost)
	if err == nil {
		t.Log(string(encrypted))
	}
}
