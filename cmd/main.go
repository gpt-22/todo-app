package main

import (
	"github.com/gtikhomiroff/todo-app"
	"github.com/gtikhomiroff/todo-app/pkg/handler"
	"github.com/gtikhomiroff/todo-app/pkg/repository"
	"github.com/gtikhomiroff/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Could not init config: %s", err)
	}

	Repository := repository.NewRepository()
	Service := service.NewService(Repository)
	Handler := handler.NewHandler(Service)

	server := new(todo_app.Server)
	if err := server.Start(viper.GetString("port"), Handler.InitRouter()); err != nil {
		log.Fatalf("Error occured while http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
