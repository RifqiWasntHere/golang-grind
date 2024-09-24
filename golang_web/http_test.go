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

// Handler with HEADER
func HeaderHandler(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, "Content Type is : "+contentType)
}

func TestHeaderHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	HeaderHandler(recorder, request)

	response := recorder.Result()
	payload, _ := io.ReadAll(response.Body)

	fmt.Println(string(payload))

}

func ReqHeaderHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("x-powered-by", "rifqicikiwir")
	fmt.Fprint(writer, "OK")
}

func TestReqHeaderHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ReqHeaderHandler(recorder, request)

	reqHeader := recorder.Header().Get("x-powered-by")
	fmt.Println(reqHeader)
}

// "POST" Form
func PostHandler(writer http.ResponseWriter, request *http.Request) {
	// Below, is the base snippet of how to fetch request body
	// err := request.ParseForm()
	// if err != nil {
	// 	log.Fatal("Parsing Error: ", err)
	// }

	// firstName := request.PostForm.Get("firstName")
	// lastName := request.PostForm.Get("lastName")

	firstName := request.PostFormValue("firstName")
	lastName := request.PostFormValue("lastName")

	fmt.Fprintf(writer, "Hey There, %s %s !", firstName, lastName)
}

func TestPostHandler(t *testing.T) {

	requestBody := strings.NewReader("firstName=Rifqi&lastName=Fadhillah")
	request := httptest.NewRequest("POST", "http://localhost:8080/", requestBody) //Somehow, "Reader" is utilized as the func to write body here
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	PostHandler(recorder, request)

	response := recorder.Result()
	payload, _ := io.ReadAll(response.Body)

	fmt.Println(string(payload))
}
