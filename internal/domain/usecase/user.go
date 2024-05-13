package usecase

import (
	"auth-service/internal/domain/dto"
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

type userUsecase struct {
	repository UserRepository
}

func NewUserUseCase(repository UserRepository) *userUsecase {
	return &userUsecase{repository: repository}
}

func (s *userUsecase) GetById(ctx context.Context, id string) (entity.User, error) {
	return s.repository.GetById(ctx, id)
}

func (s *userUsecase) GetAll(ctx context.Context) ([]entity.User, error) {
	return s.repository.GetAll(ctx)
}

func (s *userUsecase) Create(ctx context.Context, user dto.CreateUserDTO) (string, error) {
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

func (s *userUsecase) Update(ctx context.Context, user dto.UpdateUserDTO) (string, error) {

	return "", nil
}

func (s *userUsecase) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
