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
	Delete(userId, listId int) error
	Update(userId, id int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
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
		TodoItem:       NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
