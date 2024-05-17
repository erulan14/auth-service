package entity

import "time"

type User struct {
	ID          string
	Username    string
	Password    string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	IsSuperUser bool
	IsStaff     bool
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateUser struct {
	Username string
	Password string
}

type UpdateUser struct {
	Username    string
	Password    string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	IsSuperUser bool
	IsStaff     bool
	IsActive    bool
}
