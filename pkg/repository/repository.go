package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/msvetkov/notes-app/pkg/domain"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type Note interface {
	Create(note domain.Note) (int, error)
	GetAll(userId int) ([]domain.Note, error)
	GetById(userId, noteId int) (domain.Note, error)
	Delete(userId, noteId int) error
	Update(userId, noteId int, input domain.UpdateNoteInput) error
}

type User interface {
	GetById(userId int) (domain.User, error)
	Delete(userId int) error
	Update(userId int, input domain.UpdateUserInput) error
}

type Repository struct {
	Note
	Authorization
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Note:          NewNotesPostgres(db),
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
	}
}
