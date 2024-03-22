//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"time"
	"webook-server/internal/repository"
	"webook-server/internal/repository/cache"
	"webook-server/internal/repository/dao"
	"webook-server/internal/service"
	"webook-server/internal/web"
	"webook-server/internal/web/middleware"
	"webook-server/ioc"
	"webook-server/pkg/jwt"
	"webook-server/pkg/ratelimit"
)

func InitWebServer(window time.Duration, rate int) *gin.Engine {
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

		jwt.NewCacheJWTHandler,
		//wire.Value(window),
		//wire.Value(rate),
		ratelimit.NewCacheSliceWindowLimiter,
		middleware.NewAuthMiddleware,
		middleware.NewRateLimitMiddleware,

		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
	)
	return gin.Default()
}
