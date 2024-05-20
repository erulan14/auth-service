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
	Username string `db:"username" binding:"required"`
	Password string `db:"password" binding:"required"`
}

type UpdateUser struct {
	ID          string    `db:"id"`
	Username    string    `db:"username,omitempty" json:",omitempty"`
	Password    string    `db:"password" json:",omitempty"`
	FirstName   string    `db:"first_name,omitempty" json:",omitempty"`
	LastName    string    `db:"last_name" json:",omitempty"`
	Email       string    `db:"email" json:",omitempty"`
	Phone       string    `db:"phone" json:",omitempty"`
	IsActive    bool      `db:"is_active" json:",omitempty"`
	IsSuperuser bool      `db:"is_superuser" json:",omitempty"`
	IsStaff     bool      `db:"is_staff" json:",omitempty"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type DeleteUser struct {
	IsDeleted bool `db:"is_deleted"`
}
