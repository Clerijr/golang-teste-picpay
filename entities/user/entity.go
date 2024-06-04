package user

import (
	"time"

	"github.com/clerijr/teste-picpay-go/entities/user/dto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func NewUser(dto dto.NewUser) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
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
