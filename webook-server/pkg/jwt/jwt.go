package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	secret = []byte("jannan")
	exp    = 24 * time.Hour
)

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

func GenToken(userID int64, username string) (string, error) {
	c := MyClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Jannan",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
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
