package domain

import (
	"errors"
	"time"
)

type Note struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Body        string    `json:"body" db:"body" binding:"required"`
	DateCreated time.Time `json:"date_created" db:"date_created"`
	UserId      int       `json:"user_id" db:"user_id"`
}

type UpdateNoteInput struct {
	Title *string `json:"title"`
	Body  *string `json:"body"`
}

func (i UpdateNoteInput) Validate() error {
	if i.Title == nil && i.Body == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
