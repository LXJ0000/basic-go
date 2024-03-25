package model

type Model struct {
	Id       int64 `gorm:"primaryKey,autoIncrement"`
	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`
}
