package exception

import (
	"fmt"
	"go_restful_api/helper"
	"go_restful_api/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

	if notFoundError(w, r, err) {
		return
	}

	if validationErrors(w, r, err) {
		return
	}

	internalServerError(w, r, err)

}

func notFoundError(w http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		WebResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}
		fmt.Println(err)

		helper.CreateResponseBody(w, WebResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(w http.ResponseWriter, _, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		WebResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}
		fmt.Println(err)

		helper.CreateResponseBody(w, WebResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, _, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	WebResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}
	fmt.Println(err)

	helper.CreateResponseBody(w, WebResponse)
}
