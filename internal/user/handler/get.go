package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAll(ctx *gin.Context) {
	users, err := h.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting all users"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *Handler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.service.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
