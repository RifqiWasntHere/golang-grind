package main

import (
	"go_restful_api/app"
	"go_restful_api/helper"
	"go_restful_api/middleware"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Newserver(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func DBCreds() *app.DatabaseCreds {
	// load godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	// Retrieve environment variables
	return &app.DatabaseCreds{
		Credential:   os.Getenv("DB_CREDS"),
		DatabaseName: os.Getenv("DB_NAME"),
	}
}

func main() {

	server := InitializeServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
