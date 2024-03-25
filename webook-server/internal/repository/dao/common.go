package dao

import "gorm.io/gorm"

type Dao interface {
}

type CommonDao struct {
	db *gorm.DB
}

func NewCommonDao(db *gorm.DB) Dao {
	return &CommonDao{db: db}
}
