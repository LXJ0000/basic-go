//go:build wireinject

package main

import (
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

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer(window time.Duration, rate int) *gin.Engine {
	wire.Build(
		//第三方依赖
		ioc.InitDB, ioc.InitRedis,
		//dao
		dao.NewUserDao, dao.NewArticleDao,
		//cache
		cache.NewCodeCache, cache.NewUserCache, cache.NewArticleCache,
		//repository
		repository.NewUserRepository, repository.NewCodeRepository, repository.NewArticleRepository,
		//service
		ioc.InitSMSService,
		service.NewCodeService, service.NewUserService, service.NewArticleService,
		//handler
		web.NewUserHandler,
		web.NewArticleHandler,

		jwt.NewCacheJWTHandler,
		ratelimit.NewCacheSliceWindowLimiter,
		middleware.NewAuthMiddleware,
		middleware.NewRateLimitMiddleware,

		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
	)
	return gin.Default()
}
