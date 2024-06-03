package main

import (
	"log"
	"net/http"
	"os"

	"github.com/clerijr/teste-picpay-go/api/routes"
	"github.com/clerijr/teste-picpay-go/db"
	"github.com/clerijr/teste-picpay-go/user"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pgConString := os.Getenv("POSTGRE_URL")
	db.Initialize(pgConString)

	userController := user.NewController()

	controllers := routes.Controllers{
		User: userController,
	}

	router := routes.InitRoutes(controllers)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
