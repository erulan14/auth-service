package service

import (
	"auth-service/internal/user/model"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type Service interface {
	Login(req model.User) (string, error)
	VerifyToken(tokenString string) (string, error)
}

type service struct {
	secretKey string
}

func NewService(key string) Service {
	return &service{
		key,
	}
}

func (s *service) Login(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *service) VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return "", fmt.Errorf("token parse failed: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("invalid username in token")
	}

	return username, nil
}
