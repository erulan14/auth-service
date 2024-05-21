package app

import (
	v1 "auth-service/internal/api/http/v1"
	"auth-service/pkg/env"
	"auth-service/pkg/logger"
	"auth-service/pkg/pgx"
	"github.com/jmoiron/sqlx"
)

type App struct {
	env    *env.Env
	server *v1.Server
	db     *sqlx.DB
}

func NewApp() (*App, error) {
	logger.Init()
	log := logger.GetLogger()
	log.Println("logger initialized")

	a := App{}
	err := a.initDeps()

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *App) Run() error {
	defer a.db.Close()

	err := a.server.Run()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initDeps() error {
	a.env = env.NewEnv()

	var err error
	a.db, err = pgx.NewDb(a.env.DBHost, a.env.DBPort,
		a.env.DBUser, a.env.DBPass, a.env.DBName)
	if err != nil {
		return err
	}
	a.server = v1.NewServer(a.env)
	a.server.Setup(a.db)

	return nil
}
