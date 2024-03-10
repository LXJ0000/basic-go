package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//snowflake.Init("2023-01-01", 1)
	//web.InitMysql()
	//r := web.InitRouter()
	r := gin.Default()
	r.GET("ping", func(context *gin.Context) {
		context.String(200, "pong")
	})
	_ = r.Run(":8080")
}
