package controller

import (
	"encoding/json"
	"golang-library-api/service"
	"golang-library-api/web"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ControllerImpl struct {
	service service.Service
}

func NewController(service service.Service) *ControllerImpl {
	return &ControllerImpl{
		service: service,
	}
}

func (controller *ControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	requestBody := &web.CreateRequest{}
	json.NewDecoder(r.Body).Decode(requestBody)
	response := controller.service.Create(r.Context(), requestBody)

	webResponse := &web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(webResponse)

}

func (controller *ControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("bookId")
	atoi, _ := strconv.Atoi(id)
	requestBody := &web.UpdateRequest{}
	requestBody.Id = atoi
	json.NewDecoder(r.Body).Decode(requestBody)
	response := controller.service.Update(r.Context(), requestBody)

	webResponse := &web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(webResponse)

}

func (controller *ControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("bookId")
	atoi, _ := strconv.Atoi(id)

	controller.service.Delete(r.Context(), atoi)

	response := &web.WebResponse{
		Code:   204,
		Status: "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	json.NewEncoder(w).Encode(response)
}

func (controller *ControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("bookId")
	atoi, _ := strconv.Atoi(id)
	response := controller.service.FindById(r.Context(), atoi)
	webResponse := &web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	json.NewEncoder(w).Encode(webResponse)
}

func (controller *ControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	responses := controller.service.FindAll(r.Context())

	webResponse := &web.WebResponses{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(webResponse)
}
