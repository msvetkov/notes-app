package service

import (
	"github.com/msvetkov/notes-app/pkg/domain"
	"github.com/msvetkov/notes-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetById(userId int) (domain.User, error)
	Update(userId int, input domain.UpdateUserInput) error
	Delete(userId int) error
}

type Note interface {
	Create(note domain.Note) (int, error)
	GetAll(userId int) ([]domain.Note, error)
	GetById(userId int, noteId int) (domain.Note, error)
	Delete(userId int, noteId int) error
	Update(userId, noteId int, input domain.UpdateNoteInput) error
}

type Service struct {
	Note
	Authorization
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Note:          NewNotesService(repos.Note),
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
	}
}
