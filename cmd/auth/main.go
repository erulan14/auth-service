package main

import (
	"auth-service/internal/auth"
	"auth-service/internal/config"
	"log"
)

func main() {
	env := config.NewEnv()
	app := auth.New()

	err := app.Setup(env)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(env.Server)
	if err != nil {
		log.Fatal(err)
	}
}
