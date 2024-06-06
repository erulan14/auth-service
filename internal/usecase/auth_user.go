package usecase

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"auth-service/internal/auth/model"
	authService "auth-service/internal/auth/service"
	userService "auth-service/internal/user/service"
)

type UseCase interface {
	Login(ctx context.Context, request model.LoginRequest) (string, error)
	Verify(ctx context.Context, request model.VerifyRequest) error
}

type useCase struct {
	userService userService.Service
	authService authService.Service
}

func NewUseCase(userService userService.Service, authService authService.Service) UseCase {
	return &useCase{userService: userService, authService: authService}
}

func (u *useCase) Login(ctx context.Context, request model.LoginRequest) (string, error) {
	user, err := u.userService.GetByUsername(ctx, request.Username)
	if err != nil {
		return "", fmt.Errorf("error username not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", fmt.Errorf("error invalid password")
	}

	token, err := u.authService.Login(user)
	if err != nil {
		return "", fmt.Errorf("error generating token")
	}

	return token, nil
}

func (u *useCase) Verify(ctx context.Context, request model.VerifyRequest) error {
	username, err := u.authService.VerifyToken(request.Token)
	if err != nil {
		return fmt.Errorf("error verifying token: %w", err)
	}

	_, err = u.userService.GetByUsername(ctx, username)
	if err != nil {
		return fmt.Errorf("error username not found")
	}

	return nil
}
