package ioc

import (
	"github.com/gin-gonic/gin"
	"webook-server/internal/handler"
	"webook-server/internal/middleware"
)

func InitWebServer(
	middlewares []gin.HandlerFunc,
	user handler.UserHandler,
	menu handler.MenuHandler,
	blog handler.BlogInfoHandler,

) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Use(middlewares...)
	user.InitRouter(r)
	menu.InitRouter(r)
	blog.InitRouter(r)
	return r
}

func InitGinMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.CorsMiddleware(),
	}
}
