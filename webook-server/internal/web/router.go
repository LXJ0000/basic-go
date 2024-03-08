package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
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
	userGroup := r.Group("user")
	var user UserHandler
	userGroup.GET("/", user.Profile)
	userGroup.POST("/login", user.Login)
	userGroup.POST("/register", user.Register)
}
