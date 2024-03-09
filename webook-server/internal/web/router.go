package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
	"webook-server/internal/repository"
	dao2 "webook-server/internal/repository/dao"
	"webook-server/internal/service"
	"webook-server/internal/web/middleware"
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
		MaxAge: 12 * time.Hour,
	}))

	//登陆验证 https://github.com/gin-contrib/sessions
	//store := cookie.NewStore([]byte("secret")) // Cookie 存储 数据
	store, err := redis.NewStore(
		16, // 最大空闲连接
		"tcp",
		"localhost:6379",
		"",
		[]byte("8PTHDJV7izsk5f51eOmnYe5Fbrh5T17B"), // authorization key 32位
		[]byte("b1SkuScy8Mwuma90OOQMk7dwWqUMyEEF"), // encryption key 32位
	)
	if err != nil {
		panic(err)
	}
	r.Use(sessions.Sessions("my_session", store))

	initUserRouter(r)
	return r
}

func initUserRouter(r *gin.Engine) {
	dsn := "root:root@tcp(127.0.0.1:3306)/webook?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err = dao2.InitTable(db); err != nil {
		panic(err)
	}
	dao := dao2.NewUserDao(db)
	repo := repository.NewUserRepository(dao)
	svc := service.NewUserService(repo)
	user := NewUserHandler(svc)

	userGroup := r.Group("/user")
	userGroup.POST("/login", user.Login)
	userGroup.POST("/register", user.Register)

	authUserGroup := userGroup.Use(middleware.AuthMiddleware())
	authUserGroup.GET("/", user.Profile)

}
