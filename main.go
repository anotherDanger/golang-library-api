package main

import (
	"golang-library-api/controller"
	"golang-library-api/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewServer(auth *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: auth,
	}
}

func NewRouter(controller *controller.ControllerImpl) *httprouter.Router {
	router := httprouter.New()
	router.GET("/v1/library", controller.FindAll)
	router.POST("/v1/library", controller.Create)
	router.GET("/v1/library/:bookId", controller.FindById)
	router.PUT("/v1/library/:bookId", controller.Update)
	router.DELETE("/v1/library/:bookId", controller.Delete)
	return router
}

func main() {

	server := InitializeServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
