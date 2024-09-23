package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// BASIC ROUTE
func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Halo Dunia")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

// ROUTE WITH MUX
func TestMuxHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Halo Dunia")
	})
	mux.HandleFunc("/kawan", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Halo Kawan")
	})
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
