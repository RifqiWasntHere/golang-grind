package golang_web

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting User JWT")
	middleware.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("User Accessed /home")
		fmt.Fprint(w, "Welcome Home")
	})

	LogMiddleware := LogMiddleware{
		Handler: mux,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &LogMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Can't Listen : ", err)
	}
}
