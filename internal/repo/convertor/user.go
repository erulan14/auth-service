package convertor

import (
	"auth-service/internal/domain/entity"
	"auth-service/internal/repo/model"
)

func ToRepoModel(user *entity.User) *model.User {
	return &model.User{
		ID:          user.ID,
		Username:    user.Username,
		Password:    user.Password,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Phone:       user.PhoneNumber,
		IsActive:    user.IsActive,
		IsSuperuser: user.IsSuperUser,
		IsStaff:     user.IsStaff,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func ToEntityModel(user *model.User) *entity.User {
	return &entity.User{
		ID:          user.ID,
		Username:    user.Username,
		Password:    user.Password,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.Phone,
		IsActive:    user.IsActive,
		IsSuperUser: user.IsSuperuser,
		IsStaff:     user.IsStaff,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}
