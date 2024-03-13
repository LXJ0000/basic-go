package web

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"webook-server/internal/config"
	dao2 "webook-server/internal/repository/dao"
)

var db *gorm.DB

func InitMysql() {
	dsn := config.Config.DB.DSN
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err = dao2.InitTable(_db); err != nil {
		panic(err)
	}
	db = _db
}
