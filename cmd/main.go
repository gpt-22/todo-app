package main

import (
	todo "github.com/gtikhomiroff/todo-app"
	"github.com/gtikhomiroff/todo-app/pkg/handler"
	"github.com/gtikhomiroff/todo-app/pkg/repository"
	"github.com/gtikhomiroff/todo-app/pkg/service"
	"log"
)

func main() {
	Repository := repository.NewRepository()
	Service := service.NewService(Repository)
	Handler := handler.NewHandler(Service)

	server := new(todo.Server)
	if err := server.Start("8000", Handler.InitRouter()); err != nil {
		log.Fatalf("Error occured while http server: %s", err.Error())
	}
}
