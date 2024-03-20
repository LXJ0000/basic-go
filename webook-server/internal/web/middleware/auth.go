package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"webook-server/pkg/jwt"
)

type AuthMiddleware struct {
	jwt jwt.JWTHandler
	//	... 其他验证方法
}

func NewAuthMiddleware(jwt jwt.JWTHandler) *AuthMiddleware {
	return &AuthMiddleware{
		jwt: jwt,
	}
}

func (auth *AuthMiddleware) JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claim, err := auth.jwt.ParseAccessToken(ctx, parts[1])
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if claim.UserAgent != ctx.Request.UserAgent() { // 安全问题 todo 采集前端信息增强系统安全型
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err = auth.jwt.CheckSsid(ctx, claim.SSID); err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("user_id", claim.UserID)
		ctx.Set("user_name", claim.Username)
		ctx.Set("ssid", claim.SSID)
		ctx.Next()
	}
}

//func SessionAuthMiddleware() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		session := sessions.Default(ctx)
//		id := session.Get("user_id")
//		if id == nil {
//			ctx.AbortWithStatus(http.StatusUnauthorized)
//			return
//		}
//		ctx.Next()
//	}
//}
