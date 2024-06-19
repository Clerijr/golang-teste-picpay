package dto

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type UserAuth struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name" validate:"required"`
	Email string    `json:"email" validate:"required,email"`
	Exp   int64     `json:"exp,omitempty"`
	jwt.StandardClaims
}
