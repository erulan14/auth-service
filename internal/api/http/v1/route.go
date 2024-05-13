package v1

import (
	"auth-service/internal/api/http/v1/route"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Setup(db *sql.DB, gin *gin.Engine) {
	publicRouter := gin.Group("/")
	route.NewUserRouter(db, publicRouter)
}
