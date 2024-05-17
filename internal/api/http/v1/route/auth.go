package route

import (
	"auth-service/internal/api/http/v1/handler"
	"auth-service/internal/domain/service"
	"auth-service/internal/domain/usecase"
	"auth-service/internal/repo"
	"auth-service/pkg/env"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewAuth(env *env.Env, db *sqlx.DB, gin *gin.RouterGroup) {
	ur := repo.NewUser(db)
	us := service.NewUser(ur)
	au := usecase.NewAuth(env, us)
	pc := handler.NewAuth(au)
	auth := gin.Group("/auth")
	auth.POST("/login", pc.Login)
	auth.POST("/verify", pc.Verify)
}
