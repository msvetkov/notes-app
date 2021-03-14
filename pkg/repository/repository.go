package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/msvetkov/notes-app/pkg/domain"
)

type Note interface {
	Create(note domain.Note) (int, error)
	GetAll(userId int) ([]domain.Note, error)
	GetById(userId, noteId int) (domain.Note, error)
	Delete(userId, noteId int) error
	Update(userId, noteId int, input domain.UpdateNoteInput) error
}

type Repository struct {
	Note
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Note: NewNotesPostgres(db),
	}
}
