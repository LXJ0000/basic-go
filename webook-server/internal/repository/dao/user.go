package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (dao *UserDao) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.CreateAt = now
	u.UpdateAt = now
	return dao.db.WithContext(ctx).Model(&User{}).Create(&u).Error
}

type User struct {
	Id     int64 `gorm:"primaryKey,autoIncrement"`
	UserId int64 `gorm:"unique"`

	UserName string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string

	CreateAt int64
	UpdateAt int64
}

func (u *User) TableName() string {
	return `user`
}
