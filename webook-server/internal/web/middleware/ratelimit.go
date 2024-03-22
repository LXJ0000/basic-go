package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"webook-server/pkg/ratelimit"
)

type RateLimitMiddleware struct {
	cache ratelimit.RateLimit
	//	... 其他验证方法
}

func NewRateLimitMiddleware(cache ratelimit.RateLimit) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		cache: cache,
	}
}

func (l *RateLimitMiddleware) RateLimitMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isLimit, err := l.cache.Limit(ctx, l.key(ctx.ClientIP()))
		if err != nil {
			// redisOp error
			slog.Warn("RedisOp error", err.Error())
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		if isLimit {
			slog.Warn("to many req from ", ctx.ClientIP())
			ctx.AbortWithStatus(http.StatusTooManyRequests)
		}
		ctx.Next()
	}
}

func (l *RateLimitMiddleware) key(ip string) string {
	return fmt.Sprintf("ip:%s", ip)
}
