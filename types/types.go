package types

import "time"

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required,trim,min=2,max=50,alpha"`
	LastName  string `json:"lastName" validate:"required,trim,min=2,max=50,alpha"`
	Email     string `json:"email" validate:"required,trim,email,max=255"`
	Password  string `json:"password" validate:"required,min=8,max=72"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
}
