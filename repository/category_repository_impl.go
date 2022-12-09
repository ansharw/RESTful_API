package repository

import (
	"belajar-rest-api/exception"
	"belajar-rest-api/model/domain"
	"context"
	"database/sql"
)

type categoryRepository struct {
}

func NewCategoryRepository() *categoryRepository {
	return &categoryRepository{}
}

func (repo *categoryRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "select id, name From category"
	rows, err := tx.QueryContext(ctx, query)

	// helper.PanicIfError(err)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	}

	var categories []domain.Category
	for rows.Next() {
		var category domain.Category
		rows.Scan(category.GetId(), category.GetName())
		categories = append(categories, category)
	}
	return categories
}

func (repo *categoryRepository) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "insert into category(name) values(?)"
	res, err := tx.ExecContext(ctx, query, category.GetName())

	// helper.PanicIfError(err)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	}

	lastInserId, _ := res.LastInsertId()
	id := int(lastInserId)
	category.SetId(&id)

	return category
}

func (repo *categoryRepository) Update(ctx context.Context, tx *sql.Tx, category domain.Category) {
	// query := "UPDATE category SET name = NULLIF(?, '') WHERE id=?"
	query := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, category.GetName(), category.GetId())

	// helper.PanicIfError(err)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	}
}

func (repo *categoryRepository) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	query := "DELETE FROM category WHERE id=?"
	_, err := tx.ExecContext(ctx, query, category.GetId())

	// helper.PanicIfError(err)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	}
}

func (repo *categoryRepository) FindById(ctx context.Context, tx *sql.Tx, id int) domain.Category {
	query := "select id, name From category where id = ?"
	row := tx.QueryRowContext(ctx, query, id)
	var category domain.Category

	err := row.Scan(category.GetId(), category.GetName())

	// helper.PanicIfError(err)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	}

	return category
}
