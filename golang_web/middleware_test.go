package golang_web

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

// Basic Middleware
type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting User JWT")
	middleware.Handler.ServeHTTP(w, r)
}
		
// Error Handler Middleware
type ErrorHandler struct {
	Handler http.Handler
}

func (errorhandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("An Server Error Has Occured", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal Server Error")
		}
	}()

	errorhandler.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("User Accessed /home")
		fmt.Fprint(w, "Welcome Home")
	})
	mux.HandleFunc("/restricted", func(w http.ResponseWriter, r *http.Request) {
		panic("BYE BYE")
	})

	LogMiddleware := LogMiddleware{
		Handler: mux,
	}

	ErrorHandler := ErrorHandler{
		Handler: &LogMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &ErrorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Can't Listen : ", err)
	}
}
