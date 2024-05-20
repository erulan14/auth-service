package convertor

import (
	"auth-service/internal/domain/entity"
	"auth-service/internal/repo/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func ToRepoModelCreate(user entity.CreateUser) model.CreateUser {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	return model.CreateUser{
		Username: user.Username,
		Password: string(hashPassword),
	}
}

func ToRepoModelUpdate(user entity.UpdateUser) model.UpdateUser {
	return model.UpdateUser{
		Username:    user.Username,
		Password:    user.Password,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Phone:       user.Phone,
		IsActive:    user.IsActive,
		IsSuperuser: user.IsSuperUser,
		IsStaff:     user.IsStaff,
		UpdatedAt:   time.Now(),
	}
}
