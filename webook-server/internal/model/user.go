package model

import "database/sql"

type User struct {
	Model

	UserId int64 `gorm:"unique"`

	UserName sql.NullString `gorm:"unique" json:"user_name"`
	Email    sql.NullString `gorm:"unique" json:"email"`
	Phone    sql.NullString `gorm:"unique" json:"phone"`

	Password string `json:"password"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Intro    string `json:"intro"`
	WebSite  string `json:"web_site"`
}

func (u *User) TableName() string {
	return `user`
}
