package user

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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
