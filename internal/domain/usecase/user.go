package usecase

import (
	"auth-service/internal/domain/entity"
	"context"
)

type UserService interface {
	GetByUserName(ctx context.Context, username string) (entity.User, error)
	GetById(ctx context.Context, id string) (entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	Create(ctx context.Context, user entity.CreateUser) (string, error)
	Update(ctx context.Context, id string, user entity.UpdateUser) error
	Delete(ctx context.Context, id string) error
}

type user struct {
	userService UserService
}

func NewUser(userService UserService) *user {
	return &user{userService: userService}
}

func (s *user) GetById(ctx context.Context, id string) (entity.User, error) {
	repoModel, err := s.userService.GetById(ctx, id)
	return repoModel, err
}

func (s *user) GetAll(ctx context.Context) ([]entity.User, error) {
	repoModels, err := s.userService.GetAll(ctx)
	return repoModels, err
}

func (s *user) Create(ctx context.Context, user entity.CreateUser) (string, error) {
	id, err := s.userService.Create(ctx, user)
	return id, err
}

func (s *user) Update(ctx context.Context, id string, user entity.UpdateUser) error {
	return nil
}

func (s *user) Delete(ctx context.Context, id string) error {
	return s.userService.Delete(ctx, id)
}
