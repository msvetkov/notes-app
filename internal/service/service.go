package service

import "notes-app/internal/repository"

type Note interface {
}

type Service struct {
	Note
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
