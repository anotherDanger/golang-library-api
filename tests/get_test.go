package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type BookResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author,omitempty"`
}

type WebResponse struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Data   *BookResponse `json:"data,omitempty"`
}

func MockGETOK(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := &WebResponse{
		Code:   200,
		Status: "OK",
		Data: &BookResponse{
			Id:     1,
			Name:   "Programming",
			Author: "Andhika Danger",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

func MockGETFailed(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := &WebResponse{
		Code:   401,
		Status: "unauthorized",
		Data:   nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)
	json.NewEncoder(w).Encode(response)
}

func TestGETOK(t *testing.T) {
	router := httprouter.New()
	router.GET("/books", MockGETOK)
	request := httptest.NewRequest("GET", "http://localhost:8080/books", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	var WebResponses WebResponse
	json.NewDecoder(response.Body).Decode(&WebResponses)

	expected := &WebResponse{
		Code:   200,
		Status: "OK",
		Data: &BookResponse{
			Id:     1,
			Name:   "Programming",
			Author: "Andhika Danger",
		},
	}

	assert.Equal(t, expected, &WebResponses, "200 OK")
}

func TestFailed(t *testing.T) {
	router := httprouter.New()
	router.GET("/books", MockGETFailed)
	request := httptest.NewRequest("GET", "http://localhost:8080/books", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()
	var responses WebResponse
	json.NewDecoder(response.Body).Decode(&responses)
	expected := &WebResponse{
		Code:   401,
		Status: "unauthorized",
		Data:   nil,
	}

	assert.Equal(t, expected, &responses, "Not OK")
}
