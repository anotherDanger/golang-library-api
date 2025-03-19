// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"golang-library-api/controller"
	"golang-library-api/helper"
	"golang-library-api/middleware"
	"golang-library-api/repository"
	"golang-library-api/service"
	"net/http"
)

// Injectors from injector.go:

func InitializeServer() *http.Server {
	repositoryImpl := repository.NewRepository()
	db := helper.NewDb()
	serviceImpl := service.NewService(repositoryImpl, db)
	controllerImpl := controller.NewController(serviceImpl)
	router := NewRouter(controllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var BookSet = wire.NewSet(repository.NewRepository, wire.Bind(new(repository.Repository), new(*repository.RepositoryImpl)), service.NewService, wire.Bind(new(service.Service), new(*service.ServiceImpl)), controller.NewController, wire.Bind(new(controller.Controller), new(*controller.ControllerImpl)))
