package test

import (
	"database/sql"
	"go_restful_api/app"
	"go_restful_api/controller"
	"go_restful_api/middleware"
	"go_restful_api/repository"
	"go_restful_api/service"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func SetupDb(dbCreds string) *sql.DB {
	// open db connection
	db, err := sql.Open("mysql", dbCreds)
	if err != nil {
		log.Fatalln("Failed to connect to database", err)
	}

	db.SetMaxIdleConns(5)                   //Max. amount of idle dbconn
	db.SetMaxOpenConns(10)                  //Max. amount of dbconn at a same time
	db.SetConnMaxIdleTime(5 * time.Minute)  //Max. lifetime a dbconn can idle
	db.SetConnMaxLifetime(60 * time.Minute) //Max. lifetime of a dbconn

	return db
}

func SetupRouter() http.Handler {
	// load godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	// Retrieve environment variables
	dbCreds := os.Getenv("DB_CREDS")

	db := SetupDb(dbCreds)

	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// Routers
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

var BaseURI = "http://localhost:3000/api/"

func TestCreateCategorySuccess(t *testing.T) {
	router := SetupRouter()

	requestBody := strings.NewReader(`{"name": "Koplo Spanyol"}`)
	request := httptest.NewRequest(http.MethodPost, BaseURI+"categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RahasiA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}
func TestCreateCategoryFailed(t *testing.T) {

}

func TestFindByIdCategorySuccess(t *testing.T) {

}
func TestFindByIdCategoryFailed(t *testing.T) {

}

func TestFindAllCategorySuccess(t *testing.T) {

}
func TestFindAllCategoryFailed(t *testing.T) {

}

func TestUpdateCategorySuccess(t *testing.T) {

}
func TestUpdateCategoryFailed(t *testing.T) {

}

func TestDeleteCategorySuccess(t *testing.T) {

}
func TestDeleteCategoryFailed(t *testing.T) {

}

func TestUnauthorized(t *testing.T) {

}
