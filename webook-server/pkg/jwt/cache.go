package jwt

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

type CacheJWTHandler struct {
	cmd redis.Cmdable
}

func NewCacheJWTHandler(cmd redis.Cmdable) JWTHandler {
	return &CacheJWTHandler{cmd: cmd}
}

func (h *CacheJWTHandler) GenRefreshToken(ctx *gin.Context, userID int64, username string, ssid string) (string, error) {
	c := MyClaims{
		UserID:    userID,
		Username:  username,
		UserAgent: ctx.Request.UserAgent(),
		SSID:      ssid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Jannan",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshExp)),
		},
	}
	token := jwt.NewWithClaims(signingMethod, c)
	return token.SignedString(secret)
}

func (h *CacheJWTHandler) GenAccessToken(ctx *gin.Context, userID int64, username string, ssid string) (string, error) {
	c := MyClaims{
		UserID:    userID,
		Username:  username,
		UserAgent: ctx.Request.UserAgent(),
		SSID:      ssid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Jannan",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessExp)),
		},
	}
	token := jwt.NewWithClaims(signingMethod, c)
	return token.SignedString(secret)
}

func (h *CacheJWTHandler) ParseAccessToken(ctx *gin.Context, tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func (h *CacheJWTHandler) ParseRefreshToken(ctx *gin.Context, tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// DealLoginToken Get accessToken and refreshToken
func (h *CacheJWTHandler) DealLoginToken(ctx *gin.Context, userId int64, userName string) (string, string) {
	ssid := uuid.New().String()
	accessToken, _ := h.GenAccessToken(ctx, userId, userName, ssid)
	refreshToken, _ := h.GenRefreshToken(ctx, userId, userName, ssid)
	return accessToken, refreshToken
}

// CheckSsid 用户登出, token 无效
func (h *CacheJWTHandler) CheckSsid(ctx *gin.Context, ssid string) error {
	cnt, err := h.cmd.Exists(ctx, h.Key(ssid)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil // session 不存在 token 有效
		}
		return err // redisOp err
	}
	if cnt > 0 {
		return errors.New("invalid token")
	} // 无效
	return nil
}

func (h *CacheJWTHandler) InvalidateToken(ctx *gin.Context) error {
	ssid := ctx.MustGet("ssid").(string)                     // 一定不会出错 错了则 panic
	return h.cmd.Set(ctx, h.Key(ssid), "", refreshExp).Err() // key 过期时间不少于 refreshToken 过期时间
}

func (h *CacheJWTHandler) Key(ssid string) string {
	return fmt.Sprintf("user:ssid:%s", ssid)
}
