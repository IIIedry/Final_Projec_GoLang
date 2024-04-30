package main

import (
	"Application"
	"Application/internal/handlers"
	"Application/internal/repository"
	"Application/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	cfg := repository.Config{
		User: viper.GetString("POSTGRES.USER"),
		Host: viper.GetString("POSTGRES.HOST"),
		Port: viper.GetString("POSTGRES.PORT"),
		DB:   viper.GetString("POSTGRES.DB"),
		Pass: os.Getenv("DB_PASSWORD"),
	}

	db, err := repository.NewConnection(cfg)
	if err != nil {
		log.Fatalf("error initializing db connection: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)
	srv := new(Application.Server)
	err = srv.Run(viper.GetString("port"), handler.InitRoutes())
	if err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
