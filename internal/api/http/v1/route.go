package v1

import (
	"auth-service/internal/api/http/v1/route"
	"auth-service/pkg/env"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Setup(env *env.Env, db *sql.DB, gin *gin.Engine) {
	publicRouter := gin.Group("/")
	route.NewUserRouter(db, publicRouter)
}
