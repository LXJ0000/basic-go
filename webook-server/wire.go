//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"webook-server/internal/repository"
	"webook-server/internal/repository/cache"
	"webook-server/internal/repository/dao"
	"webook-server/internal/service"
	"webook-server/internal/web"
	"webook-server/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		//第三方依赖
		ioc.InitDB, ioc.InitRedis,
		//dao
		dao.NewUserDao,
		//cache
		cache.NewCodeCache, cache.NewUserCache,
		//repository
		repository.NewUserRepository, repository.NewCodeRepository,
		//service
		ioc.InitSMSService,
		service.NewCodeService, service.NewUserService,
		//handler
		web.NewUserHandler,

		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
	)
	return gin.Default()
}
