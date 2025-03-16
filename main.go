package main

import (
	"golang-library-api/controller"
	"golang-library-api/helper"
	"golang-library-api/middleware"
	"golang-library-api/repository"
	"golang-library-api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := helper.Db()
	repository := repository.NewRepository()
	service := service.NewService(repository, db)
	controller := controller.NewController(service)

	router := httprouter.New()
	router.GET("/v1/library", controller.FindAll)
	router.POST("/v1/library", controller.Create)
	router.GET("/v1/library/:bookId", controller.FindById)
	router.PUT("/v1/library/:bookId", controller.Update)
	router.DELETE("/v1/library/:bookId", controller.Delete)

	auth := &middleware.AuthMiddleware{
		Handler: router,
	}

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: auth,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
