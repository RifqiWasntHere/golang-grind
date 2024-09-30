package go_httprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

var BaseUrl = "http://localhost:8080/"

// This Params Pattern is Named : Named Parameter
func TestHttpRouter(t *testing.T) {

	router := httprouter.New()
	router.GET("/concert/:concertId/ticket/:ticketId", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		concertId, ticketId := params.ByName("concertId"), params.ByName("concertId")
		fmt.Fprintf(w, "Concert : %s, Ticket ID : %s", concertId, ticketId)
	})

	req := httptest.NewRequest("GET", BaseUrl+"concert/kickyourass/ticket/AsdDSjkSc123e", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	// Testing
	assert.Equal(t, "Concert : kickyourass, Ticket ID : kickyourass", string(body))
}

// Panic Handler
func TestPanicHandler(t *testing.T) {

	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error : ", i)
	}
	router.GET("/panic", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		panic("oopsies")
	})

	req := httptest.NewRequest("GET", BaseUrl+"panic", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Internal Server Error : oopsies", string(body))
}

// Page Not Found Handler
func TestNotFoundHandler(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Page Not Found 404")
	})

	req := httptest.NewRequest("GET", BaseUrl+"404", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	res := rec.Result()

	payload, _ := io.ReadAll(res.Body)
	assert.Equal(t, "Page Not Found 404", string(payload))
}

// Method Not Allowed Handler
func TestMethodNotAllowedHandler(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This Method Is Not allowed")
	})
	router.POST("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "hey there")
	})

	req := httptest.NewRequest("GET", BaseUrl+"", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	payload, _ := io.ReadAll(res.Body)

	assert.Equal(t, "This Method Is Not allowed", string(payload))
}
