package model

type Model struct {
	Id       int64 `gorm:"primaryKey,autoIncrement"`
	CreateAt int64
	UpdateAt int64
}
