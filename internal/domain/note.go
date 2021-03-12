package domain

import "time"

type Note struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"id"`
	Body        string    `json:"body" db:"body"`
	DateCreated time.Time `json:"date_created" db:"date_created"`
	UserId      int       `json:"user_id" db:"user_id"`
}
