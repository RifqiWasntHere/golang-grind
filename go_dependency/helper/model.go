package helper

import (
	"go_restful_api/model/domain"
	"go_restful_api/model/web"
)

// This shit is a helper to convert domain.Category struct within service implementation into web.CategoryResponse struct
func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
