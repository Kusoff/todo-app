package repository

import (
	"fmt"

	"github.com/Kusoff/todo-app"
	"github.com/jmoiron/sqlx"
)

type Auth_Postgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *Auth_Postgres {
	return &Auth_Postgres{db: db}
}

func (r *Auth_Postgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Auth_Postgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err

}
