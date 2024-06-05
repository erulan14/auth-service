package pg

import (
	"auth-service/internal/user/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	DB *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db}
}

func (s *Storage) Create(ctx context.Context, user model.User) (string, error) {
	var id string

	err := s.DB.QueryRowContext(ctx,
		`INSERT INTO "user" (username, password) VALUES ($1, $2) RETURNING id`,
		&user.Username, &user.Password).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("create user: %w", err)
	}

	return id, nil
}

func (s *Storage) Update(ctx context.Context, user model.User) error {
	_, err := s.DB.NamedExecContext(ctx, `UPDATE "user" SET 
                  username = :username, password = :password, 
                  first_name = :first_name, last_name = :last_name,
                  email = :email, phone = :phone, 
                  is_active = :is_active, is_staff = :is_staff, 
                  is_superuser=:is_superuser, updated_at=:updated_at WHERE id = :id`, &user)

	if err != nil {
		return fmt.Errorf("storage: update user, %w", err)
	}
	return nil
}

func (s *Storage) GetByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User
	err := s.DB.GetContext(ctx, &user,
		`SELECT * FROM "user" WHERE is_deleted=FALSE AND username=$1`, username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, fmt.Errorf("user not found")
		}
		return model.User{}, fmt.Errorf("storage: get user, %w", err)
	}

	return user, nil
}

func (s *Storage) GetByID(ctx context.Context, id string) (model.User, error) {
	var user model.User
	err := s.DB.GetContext(ctx, &user,
		`SELECT * FROM "user" WHERE is_deleted=FALSE AND id=$1`, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, fmt.Errorf("user not found")
		}
		return model.User{}, fmt.Errorf("storage: get user, %w", err)
	}

	return user, nil
}

func (s *Storage) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := s.DB.SelectContext(ctx, &users,
		`SELECT * FROM "user" WHERE is_deleted=false`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return make([]model.User, 0), fmt.Errorf("no users found")
		}
		return make([]model.User, 0), fmt.Errorf("storage: get users, %w", err)
	}
	return users, nil
}

func (s *Storage) Delete(ctx context.Context, id string) error {
	_, err := s.DB.ExecContext(ctx,
		`UPDATE "user" SET is_deleted = $2 WHERE id = $1`, id, true)
	if err != nil {
		return fmt.Errorf("storage: delete user, %w", err)
	}
	return nil
}
