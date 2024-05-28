package handler

import (
	"auth-service/internal/user/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Create(ctx *gin.Context) {
	var createUser model.CreateUserRequest
	if err := ctx.ShouldBind(&createUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.Create(ctx.Request.Context(), createUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}
