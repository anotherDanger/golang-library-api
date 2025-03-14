package service

import (
	"context"
	"database/sql"
	"golang-library-api/entity"
	"golang-library-api/helper"
	"golang-library-api/repository"
	"golang-library-api/web"
)

type ServiceImpl struct {
	repository *repository.RepositoryImpl
	db         *sql.DB
}

func NewService(repository *repository.RepositoryImpl, db *sql.DB) Service {
	return &ServiceImpl{
		repository: repository,
		db:         db,
	}
}

func (service *ServiceImpl) Create(ctx context.Context, request *web.CreateRequest) *web.BookResponse {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}

	entity := &entity.Domain{
		Name:   request.Name,
		Author: request.Author,
	}

	result := service.repository.Create(ctx, tx, entity)
	helper.CommitOrRollback(tx, err)
	response := helper.ToResponse(result)

	return response
}

func (service *ServiceImpl) Update(ctx context.Context, request *web.UpdateRequest) *web.BookResponse {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}

	entity := &entity.Domain{
		Id:     request.Id,
		Name:   request.Name,
		Author: request.Author,
	}

	result := service.repository.Update(ctx, tx, entity)
	helper.CommitOrRollback(tx, err)

	response := helper.ToResponse(result)

	return response
}

func (service *ServiceImpl) Delete(ctx context.Context, bookId int) {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	service.repository.Delete(ctx, tx, bookId)
	helper.CommitOrRollback(tx, err)
}

func (service *ServiceImpl) FindById(ctx context.Context, bookId int) *web.BookResponse {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}

	entity := service.repository.FindById(ctx, tx, bookId)
	helper.CommitOrRollback(tx, err)
	response := helper.ToResponse(entity)
	return response
}

func (service *ServiceImpl) FindAll(ctx context.Context) []*web.BookResponse {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}

	var result []*web.BookResponse
	entities := service.repository.FindAll(ctx, tx)
	helper.CommitOrRollback(tx, err)
	for _, entity := range entities {
		response := helper.ToResponse(entity)
		result = append(result, response)
	}

	return result
}
