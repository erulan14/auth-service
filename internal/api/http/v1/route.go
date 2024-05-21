package v1

import (
	"auth-service/internal/api/http/v1/route"
	"auth-service/pkg/env"
	"auth-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	ginlogrus "github.com/toorop/gin-logrus"
)

type Server struct {
	engine *gin.Engine
	env    *env.Env
}

func NewServer(env *env.Env) *Server {
	gin.SetMode(env.GIN_MODE)
	engine := gin.Default()
	err := engine.SetTrustedProxies(nil)
	if err != nil {
		return nil
	}

	log := logger.GetLogger()

	engine.Use(ginlogrus.Logger(log), gin.Recovery())

	return &Server{
		engine: engine,
		env:    env,
	}
}

func (s *Server) Setup(db *sqlx.DB) {
	api := s.engine.Group("/api/")
	route.NewUser(db, api)
	route.NewAuth(s.env, db, api)
}

func (s *Server) Run() error {
	return s.engine.Run(s.env.Port)
}
