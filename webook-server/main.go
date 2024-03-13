package main

import (
	"webook-server/internal/web"
	"webook-server/pkg/snowflake"
)

func main() {
	snowflake.Init("2023-01-01", 1)
	web.InitMysql()
	r := web.InitRouter()
	_ = r.Run(":8080")
}
