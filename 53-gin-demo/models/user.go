package models

import "errors"

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("empty or invalid name")
	}
	if u.Email == "" {
		return errors.New("empty or invalid email")
	}
	return nil
}
