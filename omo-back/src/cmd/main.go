package main

import (
	"log"
	omoconfig "omo-back/src/config"
	dbimpl "omo-back/src/database/impl"
	omoapp "omo-back/src/internal/app"
	"omo-back/src/internal/handler"
	serviceimpl "omo-back/src/internal/service/impl"
)

func main() {
	// PostgreSQL
	postgresService := dbimpl.NewPostgresServiceImpl()
	defer postgresService.Close()
	// User
	userRepository := dbimpl.NewUserRepositoryImpl(postgresService)
	userService := serviceimpl.NewUserServiceImpl(userRepository)

	// Handler
	handlers := omoapp.Handlers{
		AuthHandler: handler.NewAuthHandler(userService),
	}

	application := omoapp.NewApp(handlers)

	err := application.Run(":" + omoconfig.Port)
	if err != nil {
		log.Fatalf("failed to start server on port %s: %v", omoconfig.Port, err)
	}
}
