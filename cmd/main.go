package main

import (
	"log"

	"github.com/myownhatred/botAPI"
	"github.com/myownhatred/botAPI/pkg/handler"
	"github.com/myownhatred/botAPI/pkg/repository"
	"github.com/myownhatred/botAPI/pkg/service"
)

func main() {
	rep := repository.NewRepository()
	repG := repository.NewGoogleRepository()
	serv := service.NewService(rep)
	servG := service.NewGoogleService(repG)
	handlers := handler.NewHandler(serv, servG)
	srv := new(botAPI.Server)
	if err := srv.Run("6666", handlers.InitRoutes()); err != nil {
		log.Fatalf("Server run error: %s", err.Error())
	}
}
