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
			if strings.Contains(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "www.example.com")
		},
		ExposeHeaders: []string{"token"},
		MaxAge:        12 * time.Hour,
	})
}
