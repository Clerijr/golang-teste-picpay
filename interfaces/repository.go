package interfaces

import (
	"github.com/clerijr/teste-picpay-go/entities/types"
	"github.com/google/uuid"
)

type Repository interface {
	Save(dto *types.NewUser) error
	FindByID(id string) (*types.User, error)
	FindByEmail(email string) (*types.UserAuth, error)
	GetUserPassword(id uuid.UUID) (*string, error)
}
