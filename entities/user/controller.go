package user

import (
	"fmt"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("funfo")
	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("logou")
	w.WriteHeader(http.StatusOK)
}
