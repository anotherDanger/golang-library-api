package helper

import (
	"golang-library-api/entity"
	"golang-library-api/web"
)

func ToResponse(response *entity.Domain) *web.BookResponse {
	return &web.BookResponse{
		Id:     response.Id,
		Name:   response.Name,
		Author: response.Author,
	}
}
