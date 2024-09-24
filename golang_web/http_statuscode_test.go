package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func StatusCodeHandler(writer http.ResponseWriter, request *http.Request) {

	name := request.PostFormValue("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "Name is empty")
	} else {
		fmt.Fprintf(writer, "Hey There, %s !", name)
	}
}

func TestStatusCodeHandler(t *testing.T) {

	requestBody := strings.NewReader("name=Rifqi")
	request := httptest.NewRequest("POST", "http://localhost:8080/", requestBody) //Somehow, "Reader" is utilized as the func to write body here
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	StatusCodeHandler(recorder, request)

	response := recorder.Result()
	payload, _ := io.ReadAll(response.Body)
	status := response.StatusCode

	fmt.Println(string(payload), "\nStatus Code : ", status)
}
