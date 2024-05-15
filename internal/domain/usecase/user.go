package usecase

import (
	"auth-service/internal/domain/convertor"
	"auth-service/internal/domain/entity"
	"context"
	"github.com/google/uuid"
	"time"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	Create(ctx context.Context, user entity.User) error
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, id string) error
}

type user struct {
	repository UserRepository
}

func NewUser(repository UserRepository) *user {
	return &user{repository: repository}
}

func (s *user) GetById(ctx context.Context, id string) (entity.User, error) {
	return s.repository.GetById(ctx, id)
}

func (s *user) GetAll(ctx context.Context) ([]entity.User, error) {
	return s.repository.GetAll(ctx)
}

func (s *user) Create(ctx context.Context, user convertor.CreateUserDTO) (string, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	userEntity := entity.User{
		ID:          newUUID.String(),
		Username:    user.Username,
		Password:    user.Password,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.Phone,
		IsActive:    user.IsActive,
		IsStaff:     user.IsStaff,
		IsSuperUser: user.IsSuperUser,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = s.repository.Create(ctx, userEntity)
	if err != nil {
		return "", err
	}
	return newUUID.String(), nil
}

func (s *user) Update(ctx context.Context, user convertor.UpdateUserDTO) (string, error) {

	return "", nil
}

func (s *user) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
