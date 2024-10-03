package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(w http.ResponseWriter, r *http.Request, params *httprouter.Param)
	FindAll(w http.ResponseWriter, r *http.Request, params *httprouter.Param)
	FindById(w http.ResponseWriter, r *http.Request, params *httprouter.Param)
	Update(w http.ResponseWriter, r *http.Request, params *httprouter.Param)
	Delete(w http.ResponseWriter, r *http.Request, params *httprouter.Param)
}
