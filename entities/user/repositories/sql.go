package repositories

import (
	"github.com/clerijr/teste-picpay-go/entities/user"
	"github.com/clerijr/teste-picpay-go/interfaces"
	"github.com/clerijr/teste-picpay-go/types"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SQLRepo struct {
	db *sqlx.DB
}

func NewSQLRepo(db *sqlx.DB) interfaces.Repository {
	return &SQLRepo{
		db: db,
	}
}

func (r *SQLRepo) Save(dto *types.NewUser) error {

	user, _ := user.NewUser(*dto)

	_, err := r.db.NamedExec("INSERT INTO users (id, name, lastname, type, document, email, password, created_at) values (:id, :name, :lastname, :type, :document, :email, :password, :created_at)", &user)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLRepo) FindByID(id string) (*types.User, error) {
	var user types.User
	queryString := "SELECT u.id, u.name, u.lastname, u.type, u.document, u.email, u.created_at, u.updated_at, u.deleted_at, u.token, u.refresh_token FROM users u WHERE u.id=$1"

	err := r.db.Get(&user, queryString, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *SQLRepo) FindByEmail(email string) (*types.UserAuth, error) {
	var user types.UserAuth
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
