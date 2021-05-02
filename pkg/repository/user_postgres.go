package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/msvetkov/notes-app/pkg/domain"
	"strings"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetById(userId int) (domain.User, error) {
	var user domain.User

	getNoteQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&user, getNoteQuery, userId)

	return user, err
}

func (r *UserPostgres) Delete(userId int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)
	_, err := r.db.Exec(deleteQuery, userId)

	return err
}

func (r *UserPostgres) Update(userId int, input domain.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *input.Email)
		argId++
	}
	if input.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password_hash=$%d", argId))
		args = append(args, *input.Password)
		argId++
	}
	if input.Login != nil {
		setValues = append(setValues, fmt.Sprintf("login=$%d", argId))
		args = append(args, *input.Login)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", usersTable, setQuery, argId)
	args = append(args, userId)

	_, err := r.db.Exec(query, args...)
	return err
}
