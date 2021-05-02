package service

import (
	"github.com/msvetkov/notes-app/pkg/domain"
	"github.com/msvetkov/notes-app/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetById(userId int) (domain.User, error) {
	return s.repo.GetById(userId)
}

func (s *UserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s *UserService) Update(userId int, input domain.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, input)
}
