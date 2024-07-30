package main

import (
	"log"
	"net/http"
	"os"

	"github.com/clerijr/teste-picpay-go/api/routes"
	"github.com/clerijr/teste-picpay-go/entities/user"
	"github.com/clerijr/teste-picpay-go/pkg"

	"github.com/clerijr/teste-picpay-go/infra/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sqlDb, err := db.Connect()
	if err != nil {
		log.Print("Error saving user")
	}

	encoder := pkg.NewAuthEncoder()

	userRepository := user.NewSQLRepo(sqlDb)
	userController := user.NewController(userRepository, encoder)

	controllers := routes.Controllers{
		User: userController,
	}

	router := routes.InitRoutes(controllers)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
