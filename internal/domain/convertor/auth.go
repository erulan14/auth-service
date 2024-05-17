package convertor

import (
	handler_model "auth-service/internal/api/http/v1/handler/model"
	"auth-service/internal/domain/entity"
)

func ToEntityLogin(login handler_model.Login) entity.Login {
	return entity.Login{
		Username: login.Username,
		Password: login.Password,
	}
}
