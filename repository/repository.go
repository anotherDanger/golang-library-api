package repository

import (
	"context"
	"database/sql"
	"golang-library-api/entity"
)

type Repository interface {
	Create(ctx context.Context, tx *sql.Tx, book *entity.Domain) *entity.Domain
	Update(ctx context.Context, tx *sql.Tx, book *entity.Domain) *entity.Domain
	Delete(ctx context.Context, tx *sql.Tx, bookId int)
	FindById(ctx context.Context, tx *sql.Tx, bookId int) *entity.Domain
	FindAll(ctx context.Context, tx *sql.Tx) []*entity.Domain
}
