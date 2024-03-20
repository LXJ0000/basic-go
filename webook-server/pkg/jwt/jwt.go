package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	secret        = []byte("Jannan's Secret")
	accessExp     = 30 * time.Minute    // 30分钟
	refreshExp    = 15 * time.Hour * 24 // 15天
	signingMethod = jwt.SigningMethodHS256
)

type MyClaims struct {
	UserID    int64  `json:"user_id"`
	Username  string `json:"user_name"`
	UserAgent string `json:"user_agent"`
	SSID      string `json:"ssid"`
	jwt.RegisteredClaims
}

type JWTHandler interface {
	GenRefreshToken(ctx *gin.Context, userID int64, username string, ssid string) (string, error)
	GenAccessToken(ctx *gin.Context, userID int64, username string, ssid string) (string, error)
	ParseAccessToken(ctx *gin.Context, tokenString string) (*MyClaims, error)
	ParseRefreshToken(ctx *gin.Context, tokenString string) (*MyClaims, error)
	DealLoginToken(ctx *gin.Context, userId int64, userName string) (string, string)
	CheckSsid(ctx *gin.Context, ssid string) error
	InvalidateToken(ctx *gin.Context) error
}
