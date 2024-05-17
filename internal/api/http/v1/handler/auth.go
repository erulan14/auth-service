package handler

import (
	"auth-service/internal/api/http/v1/handler/model"
	"auth-service/internal/domain/convertor"
	"auth-service/internal/domain/entity"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthUseCase interface {
	Login(ctx context.Context, auth entity.Login) (string, error)
	Verify(ctx context.Context, token string) error
}

type authHandler struct {
	authUseCase AuthUseCase
}

func NewAuth(authUseCase AuthUseCase) *authHandler {
	return &authHandler{authUseCase: authUseCase}
}

func (s *authHandler) Login(ctx *gin.Context) {
	var authLogin model.Login
	if err := ctx.ShouldBind(&authLogin); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := s.authUseCase.Login(ctx.Request.Context(), convertor.ToEntityLogin(authLogin))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.ResponseToken{Token: token})
}

func (s *authHandler) Verify(ctx *gin.Context) {
	var authToken model.ResponseToken
	if err := ctx.ShouldBind(&authToken); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.authUseCase.Verify(ctx, authToken.Token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "token verified"})
}

func (s *authHandler) Register(ctx *gin.Context) {
	// TODO
}
