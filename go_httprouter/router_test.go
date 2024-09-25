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

func TestHttpRouter(t *testing.T) {

	router := httprouter.New()
	router.GET("/product/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		productId := params.ByName("id")
		fmt.Fprintf(w, "Product "+productId)
	})

	req := httptest.NewRequest("GET", "http://localhost:8080/product/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	// Testing
	assert.Equal(t, "Product 1", string(body))
}
