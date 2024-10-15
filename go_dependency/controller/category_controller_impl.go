package controller

import (
	"encoding/json"
	"go_restful_api/helper"
	"go_restful_api/model/web"
	"go_restful_api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.GetResponseBody(r, &categoryCreateRequest)

	payload := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   payload,
	}

	helper.CreateResponseBody(w, response)

}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	payload := controller.CategoryService.FindAll(r.Context())
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   payload,
	}

	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	helper.PanicIfError(err)

}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	payload := controller.CategoryService.FindById(r.Context(), categoryId)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   payload,
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.GetResponseBody(r, &categoryUpdateRequest)
	categoryUpdateRequest.Id = categoryId

	payload := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   payload,
	}

	helper.CreateResponseBody(w, response)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), categoryId)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.CreateResponseBody(w, response)
}
