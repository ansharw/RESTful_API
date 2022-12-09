package repository

import (
	"belajar-rest-api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category)
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)

	FindById(ctx context.Context, tx *sql.Tx, id int) domain.Category
}