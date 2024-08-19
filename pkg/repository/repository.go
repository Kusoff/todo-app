package repository

import (
	"github.com/Kusoff/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorithation interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorithation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorithation: NewAuthPostgres(db),
		TodoList:       NewTodoListPostgres(db),
	}
}
