package tests

import (
	"encoding/json"
	helper_test "golang-library-api/tests/helper"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func MockOKFindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	atoi, _ := strconv.Atoi(id)
	response := &helper_test.WebResponse{
		Code:   200,
		Status: "OK",
		Data: &helper_test.BookResponse{
			Id:     atoi,
			Name:   "Programming",
			Author: "Andhika Danger",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

func MockFailedFindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := &helper_test.WebResponse{
		Code:   401,
		Status: "unauthorized",
		Data:   nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)
	json.NewEncoder(w).Encode(response)
}

func TestFindByIdOK(t *testing.T) {
	router := httprouter.New()
	router.GET("/v1/library/:id", MockOKFindById)
	request := httptest.NewRequest("GET", "http://localhost:8080/v1/library/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	var WebResponses helper_test.WebResponse
	json.NewDecoder(response.Body).Decode(&WebResponses)

	expected := &helper_test.WebResponse{
		Code:   200,
		Status: "OK",
		Data: &helper_test.BookResponse{
			Id:     1,
			Name:   "Programming",
			Author: "Andhika Danger",
		},
	}

	assert.Equal(t, expected, &WebResponses, "200 OK")
}

func TestFailed(t *testing.T) {
	router := httprouter.New()
	router.GET("/v1/library", MockFailedFindById)
	request := httptest.NewRequest("GET", "http://localhost:8080/v1/library", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()
	var responses helper_test.WebResponse
	json.NewDecoder(response.Body).Decode(&responses)
	expected := &helper_test.WebResponse{
		Code:   401,
		Status: "unauthorized",
		Data:   nil,
	}

	assert.Equal(t, expected, &responses, "Not OK")
}
