package go_json

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

var BaseUrl = "http://localhost:8080/"

func TestJsonStreamingDecoder(t *testing.T) {
	router := httprouter.New()
	// Router
	router.POST("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		type payload struct {
			Name           string `json:"name"`
			UserOccupation string `json:"user_occupation"`
		}

		decoder := json.NewDecoder(r.Body)
		customerPayload := &payload{}

		_ = decoder.Decode(customerPayload)

		fmt.Fprintf(w, "Received: %s, Occupation: %s", customerPayload.Name, customerPayload.UserOccupation)
	})

	// Create a request with JSON payload
	jsonPayload := `{"name":"rifqi","user_occupation":"Yapper Handal"}`
	req := httptest.NewRequest("POST", BaseUrl, strings.NewReader(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Assert the result
	res := rec.Result()
	body, _ := io.ReadAll(res.Body)
	assert.Equal(t, "Received: rifqi, Occupation: Yapper Handal", string(body))
}

// Send JSON response using Encoder
func TestJsonStreamingEncoder(t *testing.T) {
	router := httprouter.New()
	// Router
	router.POST("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		type payload struct {
			Name           string `json:"name"`
			UserOccupation string `json:"user_occupation"`
		}

		resPayload := payload{
			Name:           "Rifqi",
			UserOccupation: "Icikiwir Tuwar Tuwir",
		}

		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)

		if err := encoder.Encode(resPayload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	// Create a request with JSON payload
	jsonPayload := `{"name":"rifqi","user_occupation":"Yapper Handal"}`
	req := httptest.NewRequest("POST", BaseUrl, strings.NewReader(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Assert the result
	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	// Print the response body (for testing purposes)
	fmt.Println(string(body))

	// Assert that the Content-Type is JSON and check the response body
	assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
	assert.JSONEq(t, `{"name":"Rifqi","user_occupation":"Icikiwir Tuwar Tuwir"}`, string(body))
}
