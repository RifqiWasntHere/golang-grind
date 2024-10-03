package main

import (
	"go_restful_api/app"
	"go_restful_api/controller"
	"go_restful_api/repository"
	"go_restful_api/service"
	"log"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
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
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
}
