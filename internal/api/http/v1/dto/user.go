package dto

import "time"

type UserDTO struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	IsSuperUser bool      `json:"is_superuser"`
	IsStaff     bool      `json:"is_staff"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateUserDTO struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
	IsSuperUser bool   `json:"is_superuser"`
	IsStaff     bool   `json:"is_staff"`
	IsActive    bool   `json:"is_active"`
}

type UpdateUserDTO struct {
	Username    string `json:"username"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	IsSuperUser bool   `json:"is_superuser"`
	IsStaff     bool   `json:"is_staff"`
	IsActive    bool   `json:"is_active"`
}
