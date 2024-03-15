package main

import (
	"webook-server/pkg/snowflake"
)

func main() {
	snowflake.Init("2023-01-01", 1)
	r := InitWebServer()
	_ = r.Run(":8080")
}
