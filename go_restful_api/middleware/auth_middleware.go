package middleware

import (
	"go_restful_api/helper"
	"go_restful_api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Api-Key") == "RahasiA" {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		WebResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "UNAUTHORIZED",
		}

		helper.CreateResponseBody(w, WebResponse)
	}

}
