package dto

import "time"

type CreateUserDTO struct {
	Username    string
	Password    string
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	IsSuperUser bool
	IsStaff     bool
	IsActive    bool
}

type UpdateUserDTO struct {
}

type UserDTO struct {
	ID          string
	Username    string
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	IsSuperUser bool
	IsStaff     bool
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
