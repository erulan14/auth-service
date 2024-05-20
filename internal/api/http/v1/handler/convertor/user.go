package convertor

import (
	"auth-service/internal/api/http/v1/handler/model"
	"auth-service/internal/domain/entity"
)

func ToHandlerModel(user entity.User) model.User {
	return model.User{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Phone:       user.Phone,
		IsActive:    user.IsActive,
		IsSuperuser: user.IsSuperUser,
		IsStaff:     user.IsStaff,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}
