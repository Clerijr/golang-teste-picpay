package user

import (
	"log"

	"github.com/clerijr/teste-picpay-go/entities/user/dto"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db  *sqlx.DB
	log log.Logger
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db:  db,
		log: *log.Default(),
	}
}

func (r *Repository) Save(dto *dto.NewUser) error {

	user, err := NewUser(*dto)
	if err != nil {
		r.log.Print("Repository: Error parsing user", err)
		return err
	}

	_, err = r.db.NamedExec("INSERT INTO users (id, name, lastname, type, document, email, password, created_at) values (:id, :name, :lastname, :type, :document, :email, :password, :created_at)", &user)
	if err != nil {
		r.log.Print("Repository: Error saving user", err)
		return err
	}

	return nil
}
