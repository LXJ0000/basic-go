package main

import (
	"webook-server/internal/web"
)

func main() {
	r := web.InitRouter()
	_ = r.Run(":8080")
}
