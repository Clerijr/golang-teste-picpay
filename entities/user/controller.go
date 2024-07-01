package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/clerijr/teste-picpay-go/entities/types"
	"github.com/clerijr/teste-picpay-go/interfaces"
	"github.com/clerijr/teste-picpay-go/pkg"
	"github.com/go-chi/chi"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	repo    interfaces.Repository
	encoder *pkg.AuthEncoder
}

func NewController(repo interfaces.Repository) *Controller {
	return &Controller{
		repo:    repo,
		encoder: pkg.NewAuthEncoder(),
	}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var user types.NewUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Controller: Error decoding body", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.repo.Save(&user)
	if err != nil {
		http.Error(w, "Controller: Error saving user", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	user, err := c.repo.FindByID(userID)
	if err != nil {
		http.Error(w, "Controller: Error geting user by id", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(&user); err != nil {
		http.Error(w, "Controller: Error encoding user", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var loginData *types.LoginUser
	var user *types.UserAuth
	var pass *string

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "Controller: Error decoding user data", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err = c.repo.FindByEmail(loginData.Email)
	if err != nil {
		http.Error(w, "Controller: Error fetching user by email", http.StatusInternalServerError)
		return
	}

	pass, err = c.repo.GetUserPassword(user.ID)
	if err != nil {
		http.Error(w, "Controller: Error fetching user data", http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(*pass), []byte(loginData.Password)); err != nil {
		http.Error(w, "Controller: Invalid Credentials", http.StatusUnauthorized)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, err := c.encoder.GenerateToken(user)
	if err != nil {
		http.Error(w, "Controller: Error generating token", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString.AccessToken})
}

func (c *Controller) Pong(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pong")

	w.WriteHeader(http.StatusOK)
}
