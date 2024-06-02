package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Lastname  string
	UType     string
	Document  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func NewUser(name, lastname, uType, document, email, password string) *User {
	return &User{
		ID:        uuid.New(),
		Name:      name,
		Lastname:  lastname,
		UType:     uType,
		Document:  document,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
}
