package main

import (
	"log"
	"net/http"
	"os"

	"github.com/clerijr/teste-picpay-go/api/routes"
	"github.com/clerijr/teste-picpay-go/entities/user"
	"github.com/clerijr/teste-picpay-go/entities/user/repositories"
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

	userRepository := repositories.NewSQLRepo(sqlDb)
	userController := user.NewController(userRepository)

	controllers := routes.Controllers{
		User: userController,
	}

	router := routes.InitRoutes(controllers)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
