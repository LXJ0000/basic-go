package ioc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"webook-server/internal/global/config"
	"webook-server/internal/repository/dao"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	if err != nil {
		panic(err)
	}
	if err = dao.InitTable(db); err != nil {
		panic(err)
	}
	return db
}
