package auth

import (
	"auth-service/internal/config"
	userHandler "auth-service/internal/user/handler"
	"auth-service/internal/user/service"
	userPG "auth-service/internal/user/storage/pg"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type App struct {
	db     *sqlx.DB
	server *gin.Engine
}

func New() *App {
	a := App{}
	return &a
}

func (a *App) Run(server config.Server) error {
	return a.server.Run(server.Port)
}

func (a *App) Setup(env *config.Env) error {
	err := a.initDb(env.Db)
	if err != nil {
		return err
	}

	err = a.initServer(env.Server)
	if err != nil {
		return err
	}
	a.initDeps(env.Server)

	return nil
}

func (a *App) initDeps(config config.Server) {
	router := a.server.Group("/" + config.Path)

	storage := userPG.NewStorage(a.db)
	service := service.NewService(storage)
	handler := userHandler.NewHandler(service)
	handler.SetupRouter(router)
}

func (a *App) initServer(config config.Server) error {
	gin.SetMode(config.GinMode)
	a.server = gin.Default()
	return nil
}

func (a *App) initDb(config config.Db) (err error) {
	a.db, err = sqlx.Connect("pgx",
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
			config.User, config.Pass, config.Host, config.Port, config.Name))
	return err
}
