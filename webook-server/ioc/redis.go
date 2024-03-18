package ioc

import (
	"github.com/redis/go-redis/v9"
	"webook-server/internal/global/config"
)

func InitRedis() redis.Cmdable {
	return redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Addr,
	})
}
