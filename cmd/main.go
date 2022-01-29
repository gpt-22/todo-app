package main

import (
	"github.com/gtikhomiroff/todo-app"
	"github.com/gtikhomiroff/todo-app/pkg/handler"
	"github.com/gtikhomiroff/todo-app/pkg/repository"
	"github.com/gtikhomiroff/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Could not init config: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("faile to initialize db: %s", err.Error())
	}

	Repository := repository.NewRepository(db)
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
