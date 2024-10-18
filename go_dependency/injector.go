//go:build wireinject
// +build wireinject

package main

import (
	"go_restful_api/app"
	"go_restful_api/controller"
	"go_restful_api/middleware"
	"go_restful_api/repository"
	"go_restful_api/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

func InitializeServer() *http.Server {
	wire.Build(
		DBCreds,
		app.NewDB,
		validator.New,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		repository.NewCategoryRepository,
		service.NewCategoryService,
		controller.NewCategoryController,
		middleware.NewAuthMiddleware,
		Newserver,
	)

	return nil
}
