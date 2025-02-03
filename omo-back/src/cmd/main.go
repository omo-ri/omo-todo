package main

import (
	"log"
	omoconfig "omo-back/src/config"
	"omo-back/src/database"
	"omo-back/src/internal/app"
)

func main() {
	if err := database.Init(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer database.Pool.Close()

	application := app.NewApp()

	err := application.Run(":" + omoconfig.Port)
	if err != nil {
		log.Fatalf("failed to start server on port %s: %v", omoconfig.Port, err)
	}
}
