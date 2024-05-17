package model

import "time"

type User struct {
	ID          string    `db:"id"`
	Username    string    `db:"username"`
	Password    string    `db:"password"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Email       string    `db:"email"`
	Phone       string    `db:"phone"`
	IsActive    bool      `db:"is_active"`
	IsSuperuser bool      `db:"is_superuser"`
	IsStaff     bool      `db:"is_staff"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	IsDeleted   bool      `db:"is_deleted"`
}

type CreateUser struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

type UpdateUser struct {
	ID          string `db:"id"`
	Username    string `db:"username"`
	Password    string `db:"password"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Email       string `db:"email"`
	Phone       string `db:"phone"`
	IsActive    bool   `db:"is_active"`
	IsSuperuser bool   `db:"is_superuser"`
	IsStaff     bool   `db:"is_staff"`
}

type DeleteUser struct {
	IsDeleted bool `db:"is_deleted"`
}
