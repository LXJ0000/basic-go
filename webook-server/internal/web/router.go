package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"webook-server/internal/repository"
	cache2 "webook-server/internal/repository/cache"
	dao2 "webook-server/internal/repository/dao"
	"webook-server/internal/service"
	"webook-server/internal/service/sms/local"
	"webook-server/internal/web/middleware"
	"webook-server/ioc"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//跨域解决方案 https://github.com/gin-contrib/cors
	r.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // 允许带 cookie
		AllowOriginFunc: func(origin string) bool {
			if strings.Contains(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "www.example.com")
		},
		ExposeHeaders: []string{"token"},
		MaxAge:        12 * time.Hour,
	}))

	initUserRouter(r)

	return r
}

func initUserRouter(r *gin.Engine) {
	db := ioc.InitDB()
	redisCmd := ioc.InitRedis()

	dao := dao2.NewUserDao(db)
	cache := cache2.NewUserCache(redisCmd)
	repo := repository.NewUserRepository(dao, cache)
	svc := service.NewUserService(repo)

	codeCache := cache2.NewCodeCache(redisCmd)
	codeRepo := repository.NewCodeRepository(codeCache)
	smsSvc := local.NewService()
	codeSvc := service.NewCodeService(codeRepo, smsSvc)

	user := NewUserHandler(svc, codeSvc)

	userGroup := r.Group("/user")
	userGroup.POST("/login", user.Login)
	userGroup.POST("/register", user.Register)

	userGroup.POST("login/sms/code", user.SendLoginSMSCode)
	userGroup.POST("login/sms/verify", user.VerifyLoginSMSCode)

	authUserGroup := userGroup.Use(middleware.JwtAuthMiddleware())
	authUserGroup.GET("/", user.Profile)
}
