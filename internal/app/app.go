package app

import (
	v1 "auth-service/internal/api/http/v1"
	"auth-service/pkg/env"
	"auth-service/pkg/pgx"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct {
	Env    *env.Env
	Router *gin.Engine
	Db     *sqlx.DB
}

func NewApp() (*App, error) {
	a := App{}
	err := a.initDeps()

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *App) Run() error {
	defer a.Db.Close()

	err := a.Router.Run(a.Env.Port)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initDeps() error {
	a.Env = env.NewEnv()

	db, err := pgx.NewDb(a.Env.DBHost, a.Env.DBPort,
		a.Env.DBUser, a.Env.DBPass, a.Env.DBName)

	if err != nil {
		return err
	}

	a.Db = db
	a.Router = gin.Default()
	v1.Setup(a.Env, a.Db, a.Router)
	return nil
}
