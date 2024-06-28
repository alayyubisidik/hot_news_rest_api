// go:build wireinject
//+build wireinject

package main

import (
	"hot_news_2/app"
	"hot_news_2/controller"
	"hot_news_2/middleware"
	"hot_news_2/repository"
	"hot_news_2/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	controller.NewUserController,
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

var articleSet = wire.NewSet(
	repository.NewArticleRepository,
	service.NewArticleService,
	controller.NewArticleController,
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		userSet,
		categorySet,
		articleSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}