package ioc

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"webook-server/internal/web"
)

func InitWebServer(middlewares []gin.HandlerFunc, user *web.UserHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Use(middlewares...)
	user.InitRouter(r)
	return r
}

func InitGinMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		//跨域解决方案 https://github.com/gin-contrib/cors
		cors.New(cors.Config{
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
		}),
	}
}
