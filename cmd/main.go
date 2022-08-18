package main

import (
	"log"

	todo "github.com/istomin10593/reminder_todo"
	"github.com/istomin10593/reminder_todo/pck/handler"
	"github.com/istomin10593/reminder_todo/pck/repository"
	"github.com/istomin10593/reminder_todo/pck/service"
)

func main() {
	repos := repository.Newepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
