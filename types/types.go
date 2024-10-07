package types

import (
	"time"
)

type UserStore interface {
	GetUserByEmail(string) (*User, error)
	CreateUser(*User) error
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID        int       `json:"id"`
	Firstname string    `json:"firstName"`
	Lastname  string    `json:"LastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserPayload struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"LastName"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}
