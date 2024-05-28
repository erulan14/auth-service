package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.service.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
