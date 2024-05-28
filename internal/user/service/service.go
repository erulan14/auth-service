package service

import (
	"auth-service/internal/user/model"
	"auth-service/internal/user/storage"
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service interface {
	Create(ctx context.Context, user model.CreateUserRequest) (string, error)
	Update(ctx context.Context, id string, user model.UpdateUserRequest) (model.UserResponse, error)
	GetByID(ctx context.Context, id string) (model.UserResponse, error)
	GetAll(ctx context.Context) ([]model.UserResponse, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) Create(ctx context.Context, req model.CreateUserRequest) (string, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := model.User{
		Username: req.Username,
		Password: string(hashPassword),
	}

	create, err := s.storage.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return create, nil
}

func (s *service) Update(ctx context.Context, id string, req model.UpdateUserRequest) (model.UserResponse, error) {
	user := model.User{
		ID:        id,
		Username:  req.Username,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		IsActive:  req.IsActive,
		IsStaff:   req.IsStaff,
		IsSuper:   req.IsSuper,
		UpdatedAt: time.Now(),
	}

	err := s.storage.Update(ctx, user)

	if err != nil {
		return model.UserResponse{}, err
	}

	userResponse := model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
		IsStaff:   user.IsStaff,
		IsSuper:   user.IsSuper,
		UpdatedAt: user.UpdatedAt,
	}

	return userResponse, nil
}

func (s *service) GetByID(ctx context.Context, id string) (model.UserResponse, error) {
	user, err := s.storage.GetByID(ctx, id)
	if err != nil {
		return model.UserResponse{}, err
	}

	userResponse := model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
		IsStaff:   user.IsStaff,
		IsSuper:   user.IsSuper,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userResponse, nil
}

func (s *service) GetAll(ctx context.Context) ([]model.UserResponse, error) {
	users, err := s.storage.GetAll(ctx)
	if err != nil {
		return make([]model.UserResponse, 0), err
	}

	usersResponse := make([]model.UserResponse, len(users))

	for index, user := range users {
		usersResponse[index] = model.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			IsActive:  user.IsActive,
			IsStaff:   user.IsStaff,
			IsSuper:   user.IsSuper,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	return usersResponse, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	err := s.storage.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
