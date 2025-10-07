package utils

import (
	"fmt"
	"main/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId uint, email string) (*string, error) {
	secretKey := config.AppConfig.JWTSecretKey
	tokenLifeTime, err := strconv.Atoi(config.AppConfig.TokenLifeTime)
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"sub":   userId,
		"email": email,
		"exp":   time.Now().Add(time.Duration(tokenLifeTime) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.AppConfig.JWTSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
