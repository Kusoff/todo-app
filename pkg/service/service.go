package service

import (
	"github.com/Kusoff/todo-app"
	"github.com/Kusoff/todo-app/pkg/repository"
)

type Authorithation interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
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
		TodoList:       NewTodoListService(repos.TodoList),
	}
}
