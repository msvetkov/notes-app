package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	server "notes-app/internal/app"
	"notes-app/internal/handler"
	"notes-app/internal/repository"
	"notes-app/internal/service"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error inizializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to inizialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app := new(server.App)
	if err := app.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurned while running http app: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
