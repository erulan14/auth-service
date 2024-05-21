package handler

import (
	"auth-service/internal/api/http/v1/handler/convertor"
	"auth-service/internal/api/http/v1/handler/model"
	domain_convertor "auth-service/internal/domain/convertor"
	"auth-service/internal/domain/entity"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserUseCase interface {
	GetById(ctx context.Context, id string) (entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	Create(ctx context.Context, user entity.CreateUser) (string, error)
	Update(ctx context.Context, id string, user entity.UpdateUser) error
	Delete(ctx context.Context, id string) error
}

type user struct {
	userUseCase UserUseCase
}

func NewUser(userUseCase UserUseCase) *user {
	return &user{userUseCase: userUseCase}
}

func (uc *user) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := uc.userUseCase.GetById(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, convertor.ToHandlerModel(user))
}

func (uc *user) GetAll(ctx *gin.Context) {
	users, err := uc.userUseCase.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handlerUsers := make([]model.User, len(users))
	for i, user := range users {
		handlerUsers[i] = convertor.ToHandlerModel(user)
	}

	ctx.JSON(http.StatusOK, handlerUsers)
}

func (uc *user) Create(ctx *gin.Context) {
	var createUser model.CreateUser
	if err := ctx.ShouldBind(&createUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uc.userUseCase.Create(ctx.Request.Context(), domain_convertor.ToEntityCreateModel(createUser))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, id)
}

func (uc *user) Update(ctx *gin.Context) {
	var updateUser model.UpdateUser
	if err := ctx.ShouldBind(&updateUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	err := uc.userUseCase.Update(ctx.Request.Context(), id, domain_convertor.ToEntityUpdateModel(updateUser))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (uc *user) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := uc.userUseCase.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
