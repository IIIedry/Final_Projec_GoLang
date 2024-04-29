package main

import (
	"Application"
	"Application/internal/handlers"
	"Application/internal/repository"
	"Application/internal/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)
	srv := new(Application.Server)
	err := srv.Run("8000", handler.InitRoutes())
	if err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
