package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func CorsMiddleware() gin.HandlerFunc {
	//跨域解决方案 https://github.com/gin-contrib/cors
	return cors.New(cors.Config{
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // 允许带 cookie
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") || strings.HasPrefix(origin, "http://127.0.0.1") {
				return true
			}
			return strings.Contains(origin, "www.example.com")
		},
		MaxAge: 12 * time.Hour,
	})
}
