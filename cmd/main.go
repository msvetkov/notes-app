package cmd

import (
	"log"
	"todo-app/internal/app"
)

func main() {
	app := new(app.App)
	if err := app.Run("8000"); err != nil {
		log.Fatalf("error occurned while running http app: %s", err)
	}
}