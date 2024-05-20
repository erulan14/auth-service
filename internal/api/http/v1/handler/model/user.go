package model

import "time"

type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	IsActive    bool      `json:"is_active"`
	IsSuperuser bool      `json:"is_superuser"`
	IsStaff     bool      `json:"is_staff"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUser struct {
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
	IsActive    bool   `json:"isActive,omitempty"`
	IsSuperuser bool   `json:"isSuperuser,omitempty"`
	IsStaff     bool   `json:"isStaff,omitempty"`
}
