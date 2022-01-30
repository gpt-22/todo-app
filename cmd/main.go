package main

import (
	"github.com/gtikhomiroff/todo-app"
	"github.com/gtikhomiroff/todo-app/pkg/handler"
	"github.com/gtikhomiroff/todo-app/pkg/repository"
	"github.com/gtikhomiroff/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Could not init config: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
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
		logrus.Fatalf("faile to initialize db: %s", err.Error())
	}

	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	Handler := handler.NewHandler(Service)

	server := new(todo_app.Server)
	if err := server.Start(viper.GetString("port"), Handler.InitRouter()); err != nil {
		logrus.Fatalf("Error occured while http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
