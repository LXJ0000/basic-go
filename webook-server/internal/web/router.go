package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
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
