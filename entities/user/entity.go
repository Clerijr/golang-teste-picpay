package user

import (
	"database/sql"
	"time"

	"github.com/clerijr/teste-picpay-go/entities/user/dto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

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

func NewUser(dto dto.NewUser) (*User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 8)
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
