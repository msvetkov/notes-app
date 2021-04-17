package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/msvetkov/notes-app/pkg/domain"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domain.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, email, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Login, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}

func (r *AuthPostgres) DeleteUser(userId int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)
	_, err := r.db.Exec(deleteQuery, userId)

	return err
}
