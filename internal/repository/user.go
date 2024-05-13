package repository

import (
	"auth-service/internal/domain/entity"
	"context"
	"database/sql"
	"log"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (s *userRepository) GetById(ctx context.Context, id string) (entity.User, error) {
	var user entity.User

	log.Println(id)
	err := s.db.QueryRowContext(ctx,
		`SELECT id, username, first_name, last_name, 
    		email, phone_number, is_superuser, is_staff, 
    	is_active, created_at, updated_at FROM "user" WHERE id = $1`, id).Scan(
		&user.ID, &user.Username, &user.FirstName, &user.LastName,
		&user.Email, &user.PhoneNumber, &user.IsSuperUser, &user.IsStaff,
		&user.IsActive, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	users := make([]entity.User, 0)
	rows, err := s.db.QueryContext(ctx, `SELECT * FROM "user"`)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName,
			&user.LastName, &user.Email, &user.PhoneNumber, &user.IsSuperUser,
			&user.IsStaff, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *userRepository) Create(ctx context.Context, user entity.User) error {
	_, err := s.db.ExecContext(ctx, `INSERT INTO 
    "user" (id, username, password, first_name, last_name, 
            email, phone_number, is_superuser, is_staff, 
            is_active, created_at, updated_at) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`,
		&user.ID, &user.Username, &user.Password, &user.FirstName,
		&user.LastName, &user.Email, &user.PhoneNumber, &user.IsSuperUser,
		&user.IsStaff, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *userRepository) Update(ctx context.Context, user entity.User) error {
	return nil
}

func (s *userRepository) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM "user" WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
