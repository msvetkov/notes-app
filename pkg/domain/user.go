package domain

import "errors"

type User struct {
	Id       int    `json:"id" db:"id"`
	Login    string `json:"login" db:"login"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password_hash"`
}

type UpdateUserInput struct {
	Login    *string `json:"login"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (i *UpdateUserInput) Validate() error {
	if i.Login == nil && i.Email == nil && i.Password == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
