package repo

import (
	"auth-service/internal/repo/model"
	"context"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	ErrUserNotFound = "User not found"
)

type user struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *user {
	return &user{db}
}

func (s *user) GetById(ctx context.Context, id string) (model.User, error) {
	var u model.User

	err := s.db.GetContext(ctx, &u, `SELECT * FROM "user" WHERE id = ?`, id)
	log.Println(err)
	if err != nil {
		return model.User{}, err
	}

	return u, nil
}

func (s *user) GetAll(ctx context.Context) ([]model.User, error) {
	users := make([]model.User, 0)
	err := s.db.SelectContext(ctx, &users, `SELECT * FROM "user"`)
	if err != nil {
		return make([]model.User, 0), err
	}
	return users, nil
}

func (s *user) Create(ctx context.Context, user model.User) error {
	_, err := s.db.NamedExecContext(ctx, `INSERT INTO "user" 
	    (username, password, first_name, last_name, email, phone, is_superuser, is_staff, is_active) 
		VALUES (:username, :password, :first_name, :last_name, 
		        :email, :phone, :is_superuser, :is_staff, :is_active)`, &user)
	if err != nil {
		return err
	}

	return nil
}

func (s *user) Update(ctx context.Context, user model.User) error {
	_, err := s.db.NamedExecContext(ctx, `UPDATE "user" SET 
                  username = :username, password = :password, 
                  first_name = :first_name, last_name = :last_name,
                  email = :email, phone = :phone, 
                  is_active = :is_active, is_staff = :is_staff, 
                  is_superuser=:is_superuser`, &user)

	if err != nil {
		return err
	}

	return nil
}

func (s *user) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM "user" WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
