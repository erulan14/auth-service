package route

import (
	"auth-service/internal/api/http/v1/controller"
	"auth-service/internal/domain/usecase"
	"auth-service/internal/repository"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(db *sql.DB, gin *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	uc := usecase.NewUserUseCase(ur)
	pc := controller.NewUserController(uc)
	group := gin.Group("/user")
	group.GET("/", pc.GetAll)
	group.GET("/:id", pc.GetById)
	group.POST("/", pc.Create)
	group.PUT("/:id", pc.Update)
	group.DELETE("/:id", pc.Delete)
}