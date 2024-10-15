package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go_restful_api/app"
	"go_restful_api/controller"
	"go_restful_api/middleware"
	"go_restful_api/model/domain"
	"go_restful_api/repository"
	"go_restful_api/service"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func SetupTestDb() *sql.DB {

	// load godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	// open db connection
	db, err := sql.Open("mysql", os.Getenv("DB_CREDS"))
	if err != nil {
		log.Fatalln("Failed to connect to database", err)
	}

	db.SetMaxIdleConns(5)                   //Max. amount of idle dbconn
	db.SetMaxOpenConns(10)                  //Max. amount of dbconn at a same time
	db.SetConnMaxIdleTime(5 * time.Minute)  //Max. lifetime a dbconn can idle
	db.SetConnMaxLifetime(60 * time.Minute) //Max. lifetime of a dbconn

	return db
}

func SetupRouter(db *sql.DB) http.Handler {
	// Retrieve environment variables

	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// Routers
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TruncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE TABLE category")
}

var BaseURI = "http://localhost:3000/api/categories"

func TestCreateCategorySuccess(t *testing.T) {
	db := SetupTestDb()
	TruncateCategory(db)
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name": "Koplo Spanyol"}`)
	request := httptest.NewRequest(http.MethodPost, BaseURI+"categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RahasiA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	// assert.Equal(t, 200, response.StatusCode) //Simplified

	// Convert response into map
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody) // map[code:200 data:map[id:3 name:Koplo Spanyo	l] status:OK]

	assert.Equal(t, 200, int(responseBody["code"].(float64))) //"code" somehow has datatype of float64
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Koplo Spanyol", responseBody["data"].(map[string]interface{})["name"])
}
func TestCreateCategoryFailed(t *testing.T) {
	db := SetupTestDb()
	TruncateCategory(db)
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, BaseURI+"categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RahasiA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	// assert.Equal(t, 200, response.StatusCode) //Simplified

	// Convert response into map
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody) // map[code:200 data:map[id:3 name:Koplo Spanyo	l] status:OK]

	assert.Equal(t, 400, int(responseBody["code"].(float64))) //"code" somehow has datatype of float64
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestFindByIdCategorySuccess(t *testing.T) {
	db := SetupTestDb()
	TruncateCategory(db)
	router := SetupRouter(db)

	tx, _ := db.Begin()
	category := repository.NewCategoryRepository().Create(context.Background(), tx, domain.Category{
		Name: "Koplo Spanyol",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, BaseURI+"/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RahasiA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	// assert.Equal(t, 200, response.StatusCode) //Simplified

	// Convert response into map
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody) // map[code:200 data:map[id:3 name:Koplo Spanyo	l] status:OK]

	assert.Equal(t, 200, int(responseBody["code"].(float64))) //"code" somehow has datatype of float64
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Koplo Spanyol", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
}
func TestFindByIdCategoryFailed(t *testing.T) {
	db := SetupTestDb()
	TruncateCategory(db)
	router := SetupRouter(db)

	tx, _ := db.Begin()
	category := repository.NewCategoryRepository().Create(context.Background(), tx, domain.Category{
		Name: "Koplo Spanyol",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, BaseURI+"/"+strconv.Itoa(category.Id+1), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RahasiA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	// assert.Equal(t, 200, response.StatusCode) //Simplified

	// Convert response into map
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody) // map[code:200 data:map[id:3 name:Koplo Spanyo	l] status:OK]

	assert.Equal(t, 404, int(responseBody["code"].(float64))) //"code" somehow has datatype of float64
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestFindAllCategorySuccess(t *testing.T) {
	db := SetupTestDb()
	TruncateCategory(db)
	router := SetupRouter(db)

	tx, _ := db.Begin()
	category1 := repository.NewCategoryRepository().Create(context.Background(), tx, domain.Category{
		Name: "Koplo Spanyol",
	})
	category2 := repository.NewCategoryRepository().Create(context.Background(), tx, domain.Category{
		Name: "Black Dangdut",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, BaseURI, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RahasiA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	// Convert response into map
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody) // map[code:200 data:map[id:3 name:Koplo Spanyo	l] status:OK]

	assert.Equal(t, 200, int(responseBody["code"].(float64))) //"code" somehow has datatype of float64
	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody)
	var categoryResponse = responseBody["data"].([]interface{})

	categoryResponse1 := categoryResponse[0].(map[string]interface{})
	categoryResponse2 := categoryResponse[1].(map[string]interface{})

	// Payload 1
	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, category1.Name, categoryResponse1["name"])

	// Payload 2
	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
	assert.Equal(t, category2.Name, categoryResponse2["name"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := SetupTestDb()
	TruncateCategory(db)
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name": "Pop Bajigur"}`)
	var categoryId int32 = 1
	request := httptest.NewRequest(http.MethodPut, BaseURI+"categories/"+strconv.Itoa(int(categoryId)), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RahasiA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	// assert.Equal(t, 200, response.StatusCode) //Simplified

	// Convert response into map
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody) // map[code:200 data:map[id:3 name:Koplo Spanyo	l] status:OK]

	assert.Equal(t, 200, int(responseBody["code"].(float64))) //"code" somehow has datatype of float64
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Pop Bajigur", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
}
func TestUpdateCategoryFailed(t *testing.T) {
	db := SetupTestDb()
	TruncateCategory(db)
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	var categoryId int32 = 1
	request := httptest.NewRequest(http.MethodPut, BaseURI+"categories/"+strconv.Itoa(int(categoryId)), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RahasiA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	// assert.Equal(t, 200, response.StatusCode) //Simplified

	// Convert response into map
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody) // map[code:200 data:map[id:3 name:Koplo Spanyo	l] status:OK]

	assert.Equal(t, 400, int(responseBody["code"].(float64))) //"code" somehow has datatype of float64
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := SetupTestDb()
	TruncateCategory(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, BaseURI+"categories/2", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RahasiA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	// assert.Equal(t, 200, response.StatusCode) //Simplified

	assert.Equal(t, "200 OK", response.Status)

}
func TestDeleteCategoryFailed(t *testing.T) {
	db := SetupTestDb()
	TruncateCategory(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, BaseURI+"categories/4", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "RahasiA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	// assert.Equal(t, 200, response.StatusCode) //Simplified

	assert.Equal(t, "404 Not Found", response.Status)
}

func TestUnauthorized(t *testing.T) {
	db := SetupTestDb()
	TruncateCategory(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, BaseURI+"categories/4", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "SalaH")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	// assert.Equal(t, 200, response.StatusCode) //Simplified

	assert.Equal(t, "401 Unauthorized", response.Status)
}
