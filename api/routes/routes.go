package routes

import (
	"github.com/clerijr/teste-picpay-go/entities/user"
	"github.com/clerijr/teste-picpay-go/pkg"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

type Controllers struct {
	User *user.Controller
}

func InitRoutes(controllers Controllers) *chi.Mux {
	c := chi.NewRouter()
	tknAuth := pkg.NewAuthEncoder().TokenAuth

	c.Get("/", controllers.User.Pong)

	c.Group(func(r chi.Router) {
		r.Post("/", controllers.User.Create)
		r.Post("/login", controllers.User.Login)
	})

	c.Group(func(r chi.Router) {

		r.Use(jwtauth.Verifier(tknAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/user/{id}", controllers.User.GetByID)
	})

	return c
}
