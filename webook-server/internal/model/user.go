package model

import "database/sql"

type User struct {
	Model

	UserId int64 `gorm:"unique"`

	UserName sql.NullString `gorm:"unique"`
	Email    sql.NullString `gorm:"unique"`
	Phone    sql.NullString `gorm:"unique"`

	Password string
	NickName string
	Avatar   string
	Intro    string
	WebSite  string
}

func (u *User) TableName() string {
	return `user`
}
