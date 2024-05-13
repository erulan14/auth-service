package controller

import (
	"auth-service/internal/api/http/v1/dto"
	user_usercase "auth-service/internal/domain/dto"
	"auth-service/internal/domain/entity"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserUseCase interface {
	GetById(ctx context.Context, id string) (entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	Create(ctx context.Context, user user_usercase.CreateUserDTO) (string, error)
	Update(ctx context.Context, user user_usercase.UpdateUserDTO) (string, error)
	Delete(ctx context.Context, id string) error
}

type userController struct {
	userUseCase UserUseCase
}

func NewUserController(userUseCase UserUseCase) *userController {
	return &userController{userUseCase: userUseCase}
}

func (uc *userController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uc.userUseCase.GetById(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *userController) GetAll(ctx *gin.Context) {
	user, err := uc.userUseCase.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *userController) Create(ctx *gin.Context) {
	var user dto.CreateUserDTO
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	usecaseDTO := user_usercase.CreateUserDTO{
		Username:    user.Username,
		Password:    user.Password,
		FirstName:   user.Firstname,
		LastName:    user.Lastname,
		Email:       user.Email,
		Phone:       user.Phone,
		IsSuperUser: user.IsSuperUser,
		IsStaff:     user.IsStaff,
		IsActive:    user.IsActive,
	}

	create, err := uc.userUseCase.Create(ctx.Request.Context(), usecaseDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, create)
}

func (uc *userController) Update(ctx *gin.Context) {

}

func (uc *userController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := uc.userUseCase.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
