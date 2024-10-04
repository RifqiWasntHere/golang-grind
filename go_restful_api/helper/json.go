package helper

import (
	"encoding/json"
	"net/http"
)

func GetResponseBody(r *http.Request, payload interface{}) {
	err := json.NewDecoder(r.Body).Decode(payload)
	PanicIfError(err)
}

func CreateResponseBody(w http.ResponseWriter, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(payload)
	PanicIfError(err)
}
