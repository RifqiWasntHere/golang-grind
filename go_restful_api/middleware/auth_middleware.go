package middleware

import (
	"go_restful_api/helper"
	"go_restful_api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func ApiKeyMiddleware(w http.ResponseWriter, r *http.Request) ApiKey {
	if r.Header.Get("X-Api-Key") == "RahasiA" {
		return ApiKey{}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	WebResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "UNAUTHORIZED",
	}

	helper.CreateResponseBody(w, WebResponse)

}
