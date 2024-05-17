package repo

import (
	"auth-service/internal/repo/model"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

const (
	ErrUserNotFound  = "User not found: "
	ErrGetByUserName = "Failed to get user by name: "
	ErrGetByIDFailed = "Failed to get user by id: "
	ErrGetAllFailed  = "Failed to get all users: "
	ErrCreateFailed  = "Failed to create user: "
	ErrUpdateFailed  = "Failed to update user: "
	ErrDeleteFailed  = "Failed to delete user: "
)

type user struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *user {
	return &user{db}
}

func (s *user) GetByUserName(ctx context.Context, username string) (model.User, error) {
	var repoModel model.User
	err := s.db.GetContext(ctx, &repoModel, `SELECT * FROM "user" WHERE username = $1`, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, errors.New(ErrUserNotFound)
		}
		return model.User{}, errors.Join(errors.New(ErrGetByUserName), err)
	}
	return repoModel, nil
}

func (s *user) GetById(ctx context.Context, id string) (model.User, error) {
	var repoModel model.User
	err := s.db.GetContext(ctx, &repoModel, `SELECT * FROM "user" WHERE id=$1 AND is_deleted=FALSE`, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, errors.New(ErrUserNotFound)
		}
		return model.User{}, errors.Join(errors.New(ErrGetByIDFailed), err)
	}

	return repoModel, nil
}

func (s *user) GetAll(ctx context.Context) ([]model.User, error) {
	var repoUsers []model.User
	err := s.db.SelectContext(ctx, &repoUsers, `SELECT * FROM "user" WHERE is_deleted=FALSE`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return make([]model.User, 0), errors.Join(errors.New(ErrUserNotFound), err)
		}
		return make([]model.User, 0), errors.Join(errors.New(ErrGetAllFailed), err)
	}
	return repoUsers, nil
}

func (s *user) Create(ctx context.Context, user model.CreateUser) (string, error) {
	var id string

	err := s.db.QueryRowContext(ctx, `INSERT INTO "user" (username, password) 
		VALUES ($1, $2) RETURNING id;`, &user.Username, &user.Password).Scan(&id)
	if err != nil {
		return "", errors.Join(errors.New(ErrCreateFailed), err)
	}

	return id, nil
}

func (s *user) Update(ctx context.Context, id string, user model.UpdateUser) error {
	user.ID = id
	_, err := s.db.NamedExecContext(ctx, `UPDATE "user" SET 
                  username = :username, password = :password, 
                  first_name = :first_name, last_name = :last_name,
                  email = :email, phone = :phone, 
                  is_active = :is_active, is_staff = :is_staff, 
                  is_superuser=:is_superuser WHERE id = :id`, &user)

	if err != nil {
		return errors.Join(errors.New(ErrUpdateFailed), err)
	}
	return nil
}

func (s *user) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `UPDATE "user" SET is_deleted = $2 WHERE id = $1`, id, true)
	if err != nil {
		return errors.Join(errors.New(ErrDeleteFailed), err)
	}
	return nil
}
