package routes

import (
	"github.com/clerijr/teste-picpay-go/user"
	"github.com/go-chi/chi"
)

type Controllers struct {
	User *user.Controller
}

func InitRoutes(controllers Controllers) *chi.Mux {
	c := chi.NewRouter()

	c.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", controllers.User.Create)
			r.Post("/login", controllers.User.Login)
		})
	})

	return c
}
