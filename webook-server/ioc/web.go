package ioc

import (
	"github.com/gin-gonic/gin"
	"webook-server/internal/web"
	"webook-server/internal/web/middleware"
)

func InitWebServer(middlewares []gin.HandlerFunc, user *web.UserHandler, auth *middleware.AuthMiddleware) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Use(middlewares...)
	user.InitRouter(r, auth)
	return r
}

func InitGinMiddlewares(rate *middleware.RateLimitMiddleware) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.CorsMiddleware(),
		rate.RateLimitMiddleware(),
	}
}
