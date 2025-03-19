package tests

import (
	"encoding/json"

	"golang-library-api/web"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func MockOKFindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := &web.WebResponses{
		Code:   200,
		Status: "OK",
		Data: []*web.BookResponse{
			{
				Id:     1,
				Name:   "Book 1",
				Author: "Andhika",
			},
			{
				Id:     2,
				Name:   "Book 2",
				Author: "Fire Demon",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

func MockFailedFindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := web.WebResponses{
		Code:   401,
		Status: "unauthorized",
		Data:   nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)
	json.NewEncoder(w).Encode(response)
}

func TestOKFindAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/v1/library", MockOKFindAll)
	request := httptest.NewRequest("GET", "http://localhost:8080/v1/library", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	var webResponse *web.WebResponses
	json.NewDecoder(response.Body).Decode(&webResponse)

	expected := &web.WebResponses{
		Code:   200,
		Status: "OK",
		Data:   webResponse.Data,
	}

	assert.Equal(t, expected, webResponse)

}

func TestFailedFindAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/v1/library", MockFailedFindAll)

	request := httptest.NewRequest("GET", "http://localhost:8080/v1/library", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	defer result.Body.Close()

	var responses *web.WebResponses
	json.NewDecoder(result.Body).Decode(&responses)

	expected := &web.WebResponses{
		Code:   401,
		Status: "unauthorized",
	}

	assert.Equal(t, expected, responses)

}
