package service

import (
	"github.com/Kusoff/todo-app"
	"github.com/Kusoff/todo-app/pkg/repository"
)

type Authorithation interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorithation
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorithation: NewAuthServise(repos.Authorithation),
	}
}