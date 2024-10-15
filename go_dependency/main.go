package main

import (
	"go_restful_api/app"
	"go_restful_api/controller"
	"go_restful_api/helper"
	"go_restful_api/middleware"
	"go_restful_api/repository"
	"go_restful_api/service"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// load godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	// Retrieve environment variables
	dbCreds := app.DatabaseCreds{
		Credential:   os.Getenv("DB_CREDS"),
		DatabaseName: os.Getenv("DB_NAME"),
	}

	db := app.NewDB(&dbCreds)
	defer db.Close()

	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// Routers
	router := app.NewRouter(categoryController)
	// Server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
