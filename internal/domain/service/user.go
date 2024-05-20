package service

import (
	"auth-service/internal/domain/convertor"
	"auth-service/internal/domain/entity"
	repo_convertor "auth-service/internal/repo/convertor"
	"auth-service/internal/repo/model"
	"context"
	"errors"
)

const (
	ErrGetByUserName = "Failed to get user by name: "
	ErrGetByIDFailed = "Failed to get user by id: "
	ErrGetAllFailed  = "Failed to get all users: "
	ErrCreateFailed  = "Failed to create user: "
	ErrUpdateFailed  = "Failed to update user: "
	ErrDeleteFailed  = "Failed to delete user: "
)

type UserRepository interface {
	GetByUserName(ctx context.Context, username string) (model.User, error)
	GetById(ctx context.Context, id string) (model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, user model.CreateUser) (string, error)
	Update(ctx context.Context, id string, user model.UpdateUser) error
	Delete(ctx context.Context, id string) error
}

type user struct {
	repository UserRepository
}

func NewUser(repository UserRepository) *user {
	return &user{repository: repository}
}

func (s *user) GetByUserName(ctx context.Context, username string) (entity.User, error) {
	repoModel, err := s.repository.GetByUserName(ctx, username)

	if err != nil {
		return entity.User{}, errors.Join(errors.New(ErrGetByUserName), err)
	}

	return convertor.ToEntityModel(repoModel), nil
}

func (s *user) GetById(ctx context.Context, id string) (entity.User, error) {
	repoModel, err := s.repository.GetById(ctx, id)

	if err != nil {
		return entity.User{}, errors.Join(errors.New(ErrGetByIDFailed), err)
	}

	return convertor.ToEntityModel(repoModel), nil
}

func (s *user) GetAll(ctx context.Context) ([]entity.User, error) {
	repoModels, err := s.repository.GetAll(ctx)

	if err != nil {
		return []entity.User{}, errors.Join(errors.New(ErrGetAllFailed), err)
	}

	entities := make([]entity.User, len(repoModels))
	for i, repoModel := range repoModels {
		entities[i] = convertor.ToEntityModel(repoModel)
	}

	return entities, nil
}

func (s *user) Create(ctx context.Context, user entity.CreateUser) (string, error) {
	id, err := s.repository.Create(ctx, repo_convertor.ToRepoModelCreate(user))
	if err != nil {
		return "", errors.Join(errors.New(ErrCreateFailed), err)
	}
	return id, nil
}

func (s *user) Update(ctx context.Context, id string, user entity.UpdateUser) error {
	return s.repository.Update(ctx, id, repo_convertor.ToRepoModelUpdate(user))
}

func (s *user) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
