package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	secret = []byte("jannan")
	exp    = 24 * time.Hour
)

type MyClaims struct {
	UserID    int64  `json:"user_id"`
	Username  string `json:"user_name"`
	UserAgent string
	jwt.RegisteredClaims
}

func GenToken(ctx *gin.Context, userID int64, username string) (string, error) {
	c := MyClaims{
		UserID:    userID,
		Username:  username,
		UserAgent: ctx.Request.UserAgent(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Jannan",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

func ParseToken(ctx *gin.Context, tokenString string) (*MyClaims, error) {
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
