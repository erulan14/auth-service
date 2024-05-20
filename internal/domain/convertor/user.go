package convertor

import (
	handler_model "auth-service/internal/api/http/v1/handler/model"
	"auth-service/internal/domain/entity"
	repo_model "auth-service/internal/repo/model"
)

func ToEntityModel(user repo_model.User) entity.User {
	return entity.User{
		ID:          user.ID,
		Username:    user.Username,
		Password:    user.Password,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Phone:       user.Phone,
		IsActive:    user.IsActive,
		IsSuperUser: user.IsSuperuser,
		IsStaff:     user.IsStaff,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func ToEntityCreateModel(user handler_model.CreateUser) entity.CreateUser {
	return entity.CreateUser{
		Username: user.Username,
		Password: user.Password,
	}
}

func ToEntityUpdateModel(user handler_model.UpdateUser) entity.UpdateUser {
	return entity.UpdateUser{
		Username:    user.Username,
		Password:    user.Password,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Phone:       user.Phone,
		IsSuperUser: user.IsSuperuser,
		IsStaff:     user.IsStaff,
		IsActive:    user.IsActive,
	}
}
