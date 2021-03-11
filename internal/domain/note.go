package domain

import "time"

type Note struct {
	Id int `json:"id" db:"id"`
	Title string `json:"title" db:"id"`
	Body string `json:"body" db:"body"`
	Time time.Time `json:"time" db:"time"`
	UserId int `json:"user_id" db:"user_id"`
}
