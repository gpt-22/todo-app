package service

import (
	app "github.com/gtikhomiroff/todo-app"
	"github.com/gtikhomiroff/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
