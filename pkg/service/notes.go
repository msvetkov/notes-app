package service

import (
	"github.com/msvetkov/notes-app/pkg/domain"
	"github.com/msvetkov/notes-app/pkg/repository"
)

type NotesService struct {
	repo repository.Note
}

func NewNotesService(repo repository.Note) *NotesService {
	return &NotesService{repo: repo}
}

func (s NotesService) Create(note domain.Note) (int, error) {
	return s.repo.Create(note)
}

func (s *NotesService) GetAll(userId int) ([]domain.Note, error) {
	return s.repo.GetAll(userId)
}

func (s *NotesService) GetById(userId int, noteId int) (domain.Note, error) {
	return s.repo.GetById(userId, noteId)
}

func (s *NotesService) Delete(userId int, noteId int) error {
	return s.repo.Delete(userId, noteId)
}

func (s *NotesService) Update(userId, noteId int, input domain.UpdateNoteInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, noteId, input)
}
