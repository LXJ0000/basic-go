package web

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"webook-server/internal/config"
	dao2 "webook-server/internal/repository/dao"
)

var db *gorm.DB
var redisClient *redis.Client

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

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Addr,
	})
	redisClient = client
}
