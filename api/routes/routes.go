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

	c.Route("/api", func(r chi.Router) {

		r.Group(func(public chi.Router) {
			public.Get("/", controllers.User.Pong)
			public.Post("/", controllers.User.Create)
			public.Post("/login", controllers.User.Login)
		})

		r.Group(func(private chi.Router) {
			private.Use(jwtauth.Verifier(tknAuth))
			private.Use(jwtauth.Authenticator)

			private.Get("/user/{id}", controllers.User.GetByID)
		})
	})

	return c
}
