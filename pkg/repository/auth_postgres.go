package repository

import (
	"fmt"

	todo "github.com/CGNihill/go-rest-api"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	qwerty := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) value ($1 $2 $3) RETURNING id")
	row := r.db.QueryRow(qwerty, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
