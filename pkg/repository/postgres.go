package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	tableUsers         = "users"
	tableTodoLists     = "todo_lists"
	tableTodoItems     = "todo_items"
	tableUserTodoLists = "user_todo_lists"
	tableTodoListItems = "todo_list_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBName, cfg.Username, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
