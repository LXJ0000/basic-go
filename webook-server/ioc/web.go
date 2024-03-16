package ioc

import (
	"github.com/gin-gonic/gin"
	"webook-server/internal/web"
	"webook-server/internal/web/middleware"
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
		middleware.CorsMiddleware(),
	}
}
