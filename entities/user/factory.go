package user

import (
	"time"

	"github.com/clerijr/teste-picpay-go/types"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(dto types.NewUser) (*types.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 8)
	if err != nil {
		return nil, err
	}

	user := types.User{
		ID:        uuid.New(),
		Name:      dto.Name,
		Lastname:  dto.Lastname,
		UType:     dto.UType,
		Document:  dto.Document,
		Email:     dto.Email,
		Password:  string(hash),
		CreatedAt: time.Now(),
	}

	return &user, nil
}
