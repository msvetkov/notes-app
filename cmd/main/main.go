package main

import (
	"log"
	server "notes-app/internal/app"
	"notes-app/internal/handler"
)

func main() {
	handlers := new(handler.Handler)

	app := new(server.App)
	if err := app.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurned while running http app: %s", err)
	}
}
