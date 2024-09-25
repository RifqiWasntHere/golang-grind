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

// This Params Pattern is Named : Named Parameter
func TestHttpRouter(t *testing.T) {

	router := httprouter.New()
	router.GET("/concert/:concertId/ticket/:ticketId", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		concertId, ticketId := params.ByName("concertId"), params.ByName("concertId")
		fmt.Fprintf(w, "Concert : %s, Ticket ID : %s", concertId, ticketId)
	})

	req := httptest.NewRequest("GET", "http://localhost:8080/concert/kickyourass/ticket/AsdDSjkSc123e", nil)
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
		fmt.Fprint(w, "Internal Server Error : ", i)
	}
	router.GET("/panic", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		panic("oopsies")
	})

	req := httptest.NewRequest("GET", "http://localhost:8080/panic", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Internal Server Error : oopsies", string(body))
}
