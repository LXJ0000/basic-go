package dao

import (
	"gorm.io/gorm"
	"webook-server/internal/model"
)

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&model.User{})
}
