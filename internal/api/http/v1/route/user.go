package route

import (
	"auth-service/internal/api/http/v1/handler"
	"auth-service/internal/domain/service"
	"auth-service/internal/domain/usecase"
	"auth-service/internal/repo"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewUser(db *sqlx.DB, gin *gin.RouterGroup) {
	ur := repo.NewUser(db)
	us := service.NewUser(ur)
	uc := usecase.NewUser(us)
	pc := handler.NewUser(uc)
	user := gin.Group("/user")
	user.GET("/", pc.GetAll)
	user.GET("/:id", pc.GetById)
	user.POST("/", pc.Create)
	user.PUT("/:id", pc.Update)
	user.DELETE("/:id", pc.Delete)
}
