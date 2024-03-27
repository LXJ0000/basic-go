package ioc

import (
	"webook-server/internal/web"
	"webook-server/internal/web/middleware"

	"github.com/gin-gonic/gin"
)

func InitWebServer(
	middlewares []gin.HandlerFunc,
	auth *middleware.AuthMiddleware,

	user *web.UserHandler,
	article *web.ArticleHandler,
) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Use(middlewares...)
	user.InitRouter(r, auth)
	article.InitRouter(r, auth)
	return r
}

func InitGinMiddlewares(rate *middleware.RateLimitMiddleware) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.CorsMiddleware(),
		rate.RateLimitMiddleware(),
	}
}
