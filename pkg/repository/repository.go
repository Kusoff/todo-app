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
	}
}
