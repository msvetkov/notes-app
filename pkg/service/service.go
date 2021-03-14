package service

import (
	"github.com/msvetkov/notes-app/pkg/domain"
	"github.com/msvetkov/notes-app/pkg/repository"
)

type Note interface {
	Create(note domain.Note) (int, error)
	GetAll(userId int) ([]domain.Note, error)
	GetById(userId int, noteId int) (domain.Note, error)
	Delete(userId int, noteId int) error
	Update(userId, noteId int, input domain.UpdateNoteInput) error
}

type Service struct {
	Note
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Note: NewNotesService(repos.Note),
	}
}
