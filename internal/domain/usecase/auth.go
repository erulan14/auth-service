package usecase

import (
	"auth-service/internal/domain/entity"
	"auth-service/pkg/env"
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	ErrInvalidUsername = "Invalid Username"
	ErrInvalidPassword = "Invalid Password"
	ErrInvalidToken    = "Invalid Token"
)

type auth struct {
	env         *env.Env
	authService UserService
}

func NewAuth(env *env.Env, authService UserService) *auth {
	return &auth{
		env:         env,
		authService: authService,
	}
}

func (s *auth) Login(ctx context.Context, user entity.Login) (string, error) {
	authUser, err := s.authService.GetByUserName(ctx, user.Username)
	if err != nil {
		return "", errors.New(ErrInvalidUsername)
	}

	err = bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New(ErrInvalidPassword)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": authUser.Username,
		"email":    authUser.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.env.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *auth) Verify(ctx context.Context, tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.env.Secret), nil
	})

	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New(ErrInvalidToken)
	}

	return nil
}
