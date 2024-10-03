package service

import (
	"context"
	"database/sql"
	"go_restful_api/helper"
	"go_restful_api/model/domain"
	"go_restful_api/model/web"
	"go_restful_api/repository"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	service.Validate.Struct(request)

	tx, err := service.DB.Begin() //Starts TX
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Create(ctx, tx, category)

	// Ini contoh dari response "Vanilla"
	res := web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

	return res

}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin() //Starts TX
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	// Ini juga bisa dijadiin helper (kata ekolokolok)
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))
	}

	return categoryResponses
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin() //Starts TX
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	// Nah ini contoh response yang pake helper
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	service.Validate.Struct(request)

	tx, err := service.DB.Begin() //Starts TX
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	/*
		Kenapa kayak gini ?, Ini penjelasannya.
		method .update (dibawah) require domain.Category sebagai parameter. dan .FindById (diatas) returns domain.Category juga.
		karena $category udah domain.Category, makanya tinggal ganti field Name dari request.Name buat dijadiin parameter .Update

		Ketimbang harus convert request (struct web.CategoryUpdateReqest) jadi domain.Category =
		category := domain.Category{
			Id: request.Id,
			Name: request.Name,
		}
			anying males banget ngetik beginian, buku catatan gakepake
	*/

	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin() //Starts TX
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	_, err = service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, categoryId)
}
