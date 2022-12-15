package repository

import (
	"belajar-rest-api/model/domain"
	"context"
	"database/sql"
)

type ProductRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
	Create(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product)
	Delete(ctx context.Context, tx *sql.Tx, product domain.Product)

	FindById(ctx context.Context, tx *sql.Tx, id int) domain.Product

	FindProductByCategoryId(ctx context.Context, tx *sql.Tx, id int) []domain.Product
	CreateProductByCategoryId(ctx context.Context, tx *sql.Tx, product domain.Product, id int) domain.Product
	UpdateProductByCategoryId(ctx context.Context, tx *sql.Tx, product domain.Product)
	DeleteProductByCategoryId(ctx context.Context, tx *sql.Tx, product domain.Product)
}
