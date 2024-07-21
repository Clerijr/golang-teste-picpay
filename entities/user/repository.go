package user

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Save(dto *NewUser) error
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*UserAuth, error)
	GetUserPassword(id uuid.UUID) (*string, error)
}

type SQLRepo struct {
	db *sqlx.DB
}

func NewSQLRepo(db *sqlx.DB) UserRepository {
	return &SQLRepo{
		db: db,
	}
}

func (r *SQLRepo) Save(dto *NewUser) error {

	user, _ := New(*dto)

	_, err := r.db.NamedExec("INSERT INTO users (id, name, lastname, type, document, email, password, created_at) values (:id, :name, :lastname, :type, :document, :email, :password, :created_at)", &user)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLRepo) FindByID(id string) (*User, error) {
	var user User
	queryString := "SELECT u.id, u.name, u.lastname, u.type, u.document, u.email, u.created_at, u.updated_at, u.deleted_at, u.token, u.refresh_token FROM users u WHERE u.id=$1"

	err := r.db.Get(&user, queryString, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *SQLRepo) FindByEmail(email string) (*UserAuth, error) {
	var user UserAuth
	queryString := "SELECT u.id, u.name, u.email FROM users u WHERE u.email=$1"

	err := r.db.Get(&user, queryString, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *SQLRepo) GetUserPassword(id uuid.UUID) (*string, error) {
	var pass string

	queryString := "SELECT u.password FROM users u WHERE u.id=$1"

	err := r.db.Get(&pass, queryString, id)
	if err != nil {
		return nil, err

	}

	return &pass, nil
}
