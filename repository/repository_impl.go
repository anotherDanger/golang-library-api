package repository

import (
	"context"
	"database/sql"
	"golang-library-api/entity"
)

type RepositoryImpl struct{}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (repository *RepositoryImpl) Create(ctx context.Context, tx *sql.Tx, book *entity.Domain) *entity.Domain {
	query := "insert into books(name, author) values(?, ?)"
	result, err := tx.ExecContext(ctx, query, book.Name, book.Author)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return &entity.Domain{
		Id:     int(id),
		Name:   book.Name,
		Author: book.Author,
	}
}

func (repository *RepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book *entity.Domain) *entity.Domain {
	queryCurrent := "select name, author from books where id = ?"
	var currentName, currentAuthor string
	current := tx.QueryRowContext(ctx, queryCurrent, book.Id)
	current.Scan(&currentName, &currentAuthor)

	if book.Name == "" {
		book.Name = currentName
	}
	if book.Author == "" {
		book.Author = currentAuthor
	}

	query := "update books set name = ?, author = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, book.Name, book.Author, book.Id)
	if err != nil {
		panic(err)
	}

	return &entity.Domain{
		Id:     book.Id,
		Name:   book.Name,
		Author: book.Author,
	}
}

func (repository *RepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, bookId int) {
	query := "delete from books where id = ?"
	_, err := tx.ExecContext(ctx, query, bookId)
	if err != nil {
		panic(err)
	}
}

func (resposiotry *RepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId int) *entity.Domain {
	query := "select * from books where id = ?"
	var name, author string
	result := tx.QueryRowContext(ctx, query, bookId)
	result.Scan(&name, &author)

	return &entity.Domain{
		Id:     bookId,
		Name:   name,
		Author: author,
	}
}

func (repository *RepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []*entity.Domain {
	query := "select * from books"
	results, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer results.Close()
	var books []*entity.Domain

	var id int
	var name, author string

	for results.Next() {
		results.Scan(&id, &name, &author)
		books = append(books, &entity.Domain{
			Id:     id,
			Name:   name,
			Author: author,
		})
	}

	return books
}
