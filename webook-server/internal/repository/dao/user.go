package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"time"

	"gorm.io/gorm"
)

var (
	ErrDuplicateEmail = errors.New("邮箱或用户名已存在")
	ErrRecordNotFound = gorm.ErrRecordNotFound
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
	err := dao.db.WithContext(ctx).Create(&u).Error
	var me *mysql.MySQLError
	if errors.As(err, &me) {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {
			return ErrDuplicateEmail
		}
	}
	return err
}

func (dao *UserDao) QueryByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Model(&User{}).Where("email=?", email).First(&u).Error
	return u, err
}

func (dao *UserDao) QueryByUserId(ctx context.Context, userId int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Model(&User{}).Where("user_id=?", userId).First(&u).Error
	return u, err
}

type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	UserId   int64  `gorm:"unique"`
	UserName string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string

	CreateAt int64
	UpdateAt int64
}

func (u *User) TableName() string {
	return `user`
}
