package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
	"webook-server/internal/repository"
	dao2 "webook-server/internal/repository/dao"
	"webook-server/internal/service"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

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

	initUserRouter(r)
	return r
}

func initUserRouter(r *gin.Engine) {
	dsn := "root:root@tcp(127.0.0.1:33060)/webook?charset=utf8mb4&parseTime=True&loc=Local"
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

	userGroup := r.Group("user")
	userGroup.GET("/", user.Profile)
	userGroup.POST("/login", user.Login)
	userGroup.POST("/register", user.Register)
}
