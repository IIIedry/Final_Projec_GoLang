package main

import (
	"Application"
	"Application/internal/handlers"
	"Application/internal/repository"
	"Application/internal/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)
	srv := new(Application.Server)
	err := srv.Run(viper.GetString("port"), handler.InitRoutes())
	if err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
