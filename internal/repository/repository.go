package repository

import "github.com/jmoiron/sqlx"

type Note interface {
}

type Repository struct {
	Note
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
