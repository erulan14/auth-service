package handler

import (
	"auth-service/internal/auth/model"
	"auth-service/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase usecase.UseCase
}

func NewHandler(useCase usecase.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Register(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	auth.POST("/login", h.login)
	auth.POST("/verify", h.verify)
}

func (h *Handler) login(ctx *gin.Context) {
	var loginRequest model.LoginRequest
	if err := ctx.ShouldBind(&loginRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.useCase.Login(ctx.Request.Context(), loginRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) verify(ctx *gin.Context) {
	var verifyRequest model.VerifyRequest
	if err := ctx.ShouldBind(&verifyRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.Verify(ctx.Request.Context(), verifyRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Token verified successfully"})
}
