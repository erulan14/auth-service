package v1

import (
	"auth-service/internal/api/http/v1/route"
	"auth-service/pkg/env"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Setup(env *env.Env, db *sqlx.DB, gin *gin.Engine) {
	publicRouter := gin.Group("/")
	route.NewUser(db, publicRouter)
	route.NewAuth(env, db, publicRouter)
}
