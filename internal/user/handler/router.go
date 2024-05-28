package handler

import (
	"auth-service/internal/user/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SetupRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.POST("/", h.Create)
	user.PUT("/:id/", h.Update)
	user.GET("/", h.GetAll)
	user.GET("/:id/", h.GetByID)
	user.DELETE("/:id/", h.Delete)
}
