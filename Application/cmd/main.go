package main

import (
	"Application"
	"Application/internal/handlers"
	"log"
)

func main() {
	handlers := new(handlers.Handler)
	srv := new(Application.Server)
	err := srv.Run("8000", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
