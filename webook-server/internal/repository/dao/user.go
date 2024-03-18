package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
	"webook-server/internal/model"
)

type UserDao interface {
	Insert(ctx context.Context, u model.User) error
	FindByPhone(ctx context.Context, phone string) (model.User, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
	FindByUserId(ctx context.Context, userId int64) (model.User, error)
}

type UserDaoByGorm struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &UserDaoByGorm{
		db: db,
	}
}

func (dao *UserDaoByGorm) Insert(ctx context.Context, u model.User) error {
	//todo cache
	now := time.Now().UnixMilli()
	u.CreateAt = now
	u.UpdateAt = now
	return dao.db.WithContext(ctx).Create(&u).Error
}

func (dao *UserDaoByGorm) FindByPhone(ctx context.Context, phone string) (model.User, error) {
	var u model.User
	err := dao.db.WithContext(ctx).Model(&model.User{}).Where("phone=?", phone).First(&u).Error
	return u, err
}

func (dao *UserDaoByGorm) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var u model.User
	err := dao.db.WithContext(ctx).Model(&model.User{}).Where("email=?", email).First(&u).Error
	return u, err
}

func (dao *UserDaoByGorm) FindByUserId(ctx context.Context, userId int64) (model.User, error) {
	var u model.User
	err := dao.db.WithContext(ctx).Model(&model.User{}).Where("user_id=?", userId).First(&u).Error
	return u, err
}
