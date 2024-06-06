package main

import (
	"auth-service/internal/app"
	"auth-service/internal/config"
	"log"
)

func main() {
	env := config.NewEnv()
	application := app.New()

	err := application.Setup(env)
	if err != nil {
		log.Fatal(err)
	}

	err = application.Run(env.Server)
	if err != nil {
		log.Fatal(err)
	}
}
