package repository

import (
	"belajar-rest-api/exception"
	"belajar-rest-api/model/domain"
	"context"
	"database/sql"
)

type productRepository struct {
}

func NewProductRepository() *productRepository {
	return &productRepository{}
}

// "/categories"

func (repo *productRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	query := "select id, name from product"
	rows, err := tx.QueryContext(ctx, query)

	if err != nil {
		panic(exception.NewNotFoundError(err))
	}

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		rows.Scan(product.GetId(), product.GetName())
		products = append(products, product)
	}
	return products
}

func (repo *productRepository) Create(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	query := "insert into product(name, category_id) values(?,?)"
	res, err := tx.ExecContext(ctx, query, product.GetName(), product.GetCategryId())

	if err != nil {
		panic(exception.NewNotFoundError(err))
	}

	lastInsertId, _ := res.LastInsertId()
	id := int(lastInsertId)
	product.SetId(&id)

	return product
}

func (repo *productRepository) Update(ctx context.Context, tx *sql.Tx, product domain.Product) {
	query := "update product set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, product.GetName(), product.GetId())

	if err != nil {
		panic(exception.NewNotFoundError(err))
	}
}

func (repo *productRepository) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	query := "delete from product where id = ?"
	_, err := tx.ExecContext(ctx, query, product.GetId())

	if err != nil {
		panic(exception.NewNotFoundError(err))
	}
}

func (repo *productRepository) FindById(ctx context.Context, tx *sql.Tx, id int) domain.Product {
	query := "select id, name, category_id from product where id = ?"
	row := tx.QueryRowContext(ctx, query, id)
	var product domain.Product

	err := row.Scan(product.GetId(), product.GetName(), product.GetCategryId())

	if err != nil {
		panic(exception.NewNotFoundError(err))
	}

	return product
}

func (repo *productRepository) FindProductByCategoryId(ctx context.Context, tx *sql.Tx, id int) []domain.Product {
	query := "SELECT id, name FROM product WHERE category_id = ?"
	rows, err := tx.QueryContext(ctx, query, id)

	if err != nil {
		panic(exception.NewInternalServerError(err))
	}

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		rows.Scan(product.GetId(), product.GetName())
		products = append(products, product)
	}

	return products
}

func (repo *productRepository) CreateProductByCategoryId(ctx context.Context, tx *sql.Tx, product domain.Product, id int) domain.Product {
	query := "insert into product(name, category_id) values(?,?)"
	res, err := tx.ExecContext(ctx, query, product.GetName(), id)

	if err != nil {
		panic(exception.NewNotFoundError(err))
	}

	lastInsertId, _ := res.LastInsertId()
	idp := int(lastInsertId)
	product.SetId(&idp)
	product.SetCategoryId(&id)

	return product
}

func (repo *productRepository) UpdateProductByCategoryId(ctx context.Context, tx *sql.Tx, product domain.Product) {
	query := "update product set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, product.GetName(), product.GetId())

	if err != nil {
		panic(exception.NewNotFoundError(err))
	}
}

func (repo *productRepository) DeleteProductByCategoryId(ctx context.Context, tx *sql.Tx, product domain.Product) {
	query := "delete from product where id = ?"
	_, err := tx.ExecContext(ctx, query, product.GetId())

	if err != nil {
		panic(exception.NewNotFoundError(err))
	}
}