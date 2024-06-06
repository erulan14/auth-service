package handler

import (
	"auth-service/internal/user/model"
	"auth-service/internal/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.POST("/", h.create)
	user.PUT("/:id/", h.update)
	user.GET("/", h.getAll)
	user.GET("/:id/", h.getByID)
	user.DELETE("/:id/", h.delete)
}

func (h *Handler) create(ctx *gin.Context) {
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

func (h *Handler) update(ctx *gin.Context) {
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

func (h *Handler) getAll(ctx *gin.Context) {
	users, err := h.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting all users"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *Handler) getByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.service.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.service.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
