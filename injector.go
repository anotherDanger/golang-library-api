//go:build wireinject
// +build wireinject

package main

import (
	"golang-library-api/controller"
	"golang-library-api/helper"
	"golang-library-api/middleware"
	"golang-library-api/repository"
	"golang-library-api/service"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var BookSet = wire.NewSet(
	repository.NewRepository, wire.Bind(new(repository.Repository), new(*repository.RepositoryImpl)),
	service.NewService, wire.Bind(new(service.Service), new(*service.ServiceImpl)),
	controller.NewController, wire.Bind(new(controller.Controller), new(*controller.ControllerImpl)),
)

func InitializeServer() *http.Server {
	wire.Build(
		BookSet,
		helper.NewDb,
		NewRouter, wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
