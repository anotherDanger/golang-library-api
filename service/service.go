package service

import (
	"context"
	"golang-library-api/web"
)

type Service interface {
	Create(ctx context.Context, request *web.CreateRequest) *web.BookResponse
	Update(ctx context.Context, request *web.UpdateRequest) *web.BookResponse
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) *web.BookResponse
	FindAll(ctx context.Context) []*web.BookResponse
}
