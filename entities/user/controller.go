package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/clerijr/teste-picpay-go/entities/user/dto"
)

type Controller struct {
	repo *Repository
	log  log.Logger
}

func NewController(repo *Repository) *Controller {
	return &Controller{
		repo: repo,
		log:  *log.Default(),
	}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {

	var user dto.NewUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		c.log.Print("Controller: Error decoding body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.repo.Save(&user)
	if err != nil {
		c.log.Print("Controller: Error saving user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("logou")
	w.WriteHeader(http.StatusOK)
}
