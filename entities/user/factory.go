package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidName     = errors.New("invalid name")
	ErrInvalidDocument = errors.New("invalid document")
	ErrInvalidUserType = errors.New("invalid user type, can only be 'fisica' or 'juridica'")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
)

func New(user NewUser) (*User, error) {

	if user.Name == "" {
		return nil, ErrInvalidName
	}

	if user.Document == "" || len(user.Document) > 14 {
		return nil, ErrInvalidDocument
	}

	if user.UType == "" || user.UType != "fisica" && user.UType != "juridica" {
		return nil, ErrInvalidUserType
	}

	if user.Email == "" { //add a better validation later
		return nil, ErrInvalidEmail
	}

	if user.Password == "" {
		return nil, ErrInvalidPassword
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return nil, ErrInvalidPassword
	}

	return &User{
		ID:        uuid.New(),
		Name:      user.Name,
		Lastname:  user.Lastname,
		UType:     user.UType,
		Document:  user.Document,
		Email:     user.Email,
		Password:  string(hash),
		CreatedAt: time.Now(),
	}, nil
}
