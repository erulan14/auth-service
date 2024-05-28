package storage

import (
	"auth-service/internal/user/model"
	"context"
)

type Storage interface {
	GetByID(ctx context.Context, id string) (model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, model model.User) (string, error)
	Update(ctx context.Context, model model.User) error
	Delete(ctx context.Context, id string) error
}
