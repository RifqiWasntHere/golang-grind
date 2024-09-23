package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Basic Route
func RouteHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Route Loaded !")
}

// HTTPTEST
func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/Hello", nil)
	recorder := httptest.NewRecorder()

	RouteHandler(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// Route + Request Parameter
func RouteWithParamsHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello, Anon")
	} else {
		fmt.Fprintf(writer, "Hello, %s", name)
	}
}

func TestRouteWithParams(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/home?name=Rifqi", nil)
	recorder := httptest.NewRecorder()

	RouteWithParamsHandler(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// And If There's Multiple Value Within a Parameter, Do This :
func MultiValuedParams(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query() //Hasil dari Query itu berupa string slice
	payload := query["name"]
	fmt.Fprint(writer, strings.Join(payload, " "))
}
func TestMultiValuedParams(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/home?name=Rifqi&name=Fadhillah", nil)
	recorder := httptest.NewRecorder()

	MultiValuedParams(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
