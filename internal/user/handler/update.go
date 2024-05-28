package handler

import (
	"auth-service/internal/user/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var updateUser model.UpdateUserRequest
	if err := ctx.ShouldBind(&updateUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Update(ctx.Request.Context(), id, updateUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
