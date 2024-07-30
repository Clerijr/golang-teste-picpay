package user

import (
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserAuthToken struct {
	AccessToken string `json:"access_token"`
}

type UserAuth struct {
	ID    uuid.UUID `db:"id"`
	Name  string    `json:"name" validate:"required"`
	Email string    `json:"email" validate:"required,email"`
	Exp   int64     `json:"exp,omitempty"`
	jwt.StandardClaims
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Lastname string `json:"lastname" validate:"required"`
	UType    string `json:"type" validate:"required"`
	Document string `json:"document" validate:"required,min=11,max=14"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID           uuid.UUID      `db:"id"`
	Name         string         `db:"name"`
	Lastname     string         `db:"lastname"`
	UType        string         `db:"type"`
	Document     string         `db:"document"`
	Email        string         `db:"email"`
	Password     string         `db:"password"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
	DeletedAt    sql.NullTime   `db:"deleted_at"`
	Token        sql.NullString `json:"token,omitempty" db:"token"`
	RefreshToken sql.NullString `json:"refresh_token,omitempty" db:"refresh_token"`
}

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
