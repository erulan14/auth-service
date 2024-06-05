package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"auth-service/internal/config"
	userH "auth-service/internal/user/handler"
	userS "auth-service/internal/user/service"
	userPG "auth-service/internal/user/storage/pg"

	authH "auth-service/internal/auth/handler"
	authS "auth-service/internal/auth/service"
	authU "auth-service/internal/usecase"
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

	userStorage := userPG.NewStorage(a.db)
	userService := userS.NewService(userStorage)
	userHandler := userH.NewHandler(userService)
	userHandler.Register(router)

	authService := authS.NewService(config.Secret)
	authUseCase := authU.NewUseCase(userService, authService)
	authHandler := authH.NewHandler(authUseCase)
	authHandler.Register(router)
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
