package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/harshit14100/go-todo/config"
)

func getJwtKey() []byte {
	return []byte(config.GetEnv("JWT_SECRET", "default_fallback_secret"))
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJwtKey())
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return getJwtKey(), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
