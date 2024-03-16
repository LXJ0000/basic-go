package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrDuplicate      = errors.New("邮箱或用户名已存在")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type UserDao interface {
	Insert(ctx context.Context, u User) error
	FindByPhone(ctx context.Context, phone string) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	FindByUserId(ctx context.Context, userId int64) (User, error)
}

type UserDaoByGorm struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &UserDaoByGorm{
		db: db,
	}
}

func (dao *UserDaoByGorm) Insert(ctx context.Context, u User) error {
	//todo cache
	now := time.Now().UnixMilli()
	u.CreateAt = now
	u.UpdateAt = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	var me *mysql.MySQLError
	if errors.As(err, &me) {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {
			return ErrDuplicate
		}
	}
	return err
}

func (dao *UserDaoByGorm) FindByPhone(ctx context.Context, phone string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Model(&User{}).Where("phone=?", phone).First(&u).Error
	return u, err
}

func (dao *UserDaoByGorm) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Model(&User{}).Where("email=?", email).First(&u).Error
	return u, err
}

func (dao *UserDaoByGorm) FindByUserId(ctx context.Context, userId int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Model(&User{}).Where("user_id=?", userId).First(&u).Error
	return u, err
}

type User struct {
	Id     int64 `gorm:"primaryKey,autoIncrement"`
	UserId int64 `gorm:"unique"`

	UserName sql.NullString `gorm:"unique"`
	Email    sql.NullString `gorm:"unique"`
	Phone    sql.NullString `gorm:"unique"`

	Password string

	CreateAt int64
	UpdateAt int64
}

func (u *User) TableName() string {
	return `user`
}
