package service

import (
	"belajar-rest-api/helper"
	"belajar-rest-api/model/domain"
	"belajar-rest-api/model/request"
	"belajar-rest-api/model/response"
	"belajar-rest-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type productService struct {
	db                *sql.DB
	productRepository repository.ProductRepository
	// nambahin validate
	validate *validator.Validate
}

func NewProductService(db *sql.DB, productRepository repository.ProductRepository, validate *validator.Validate) *productService {
	return &productService{db, productRepository, validate}
}

func (service *productService) FindAll(ctx context.Context) []response.ResponseProduct {
	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.productRepository.FindAll(ctx, tx)

	responseCategories := []response.ResponseProduct{}
	for _, v := range categories {
		responseCategories = append(responseCategories, v.ToResponseProduct())
	}

	return responseCategories
}

func (service *productService) Create(ctx context.Context, request request.RequestCreateProduct) response.ResponseProduct {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product := domain.Product{}
	product.SetName(&request.Name)
	product.SetCategoryId(&request.CategoryId)
	product = service.productRepository.Create(ctx, tx, product)

	return product.ToResponseProduct()
}

func (service *productService) Update(ctx context.Context, request request.RequestUpdateProduct) {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product := service.productRepository.FindById(ctx, tx, request.Id)
	product.SetName(&request.Name)

	service.productRepository.Update(ctx, tx, product)
}

func (service *productService) Delete(ctx context.Context, request request.RequestDeleteProduct) {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product := service.productRepository.FindById(ctx, tx, request.Id)

	service.productRepository.Delete(ctx, tx, product)
}


func (service *productService) FindById(ctx context.Context, id int) response.ResponseProduct {
	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product := service.productRepository.FindById(ctx, tx, id)

	return product.ToResponseProduct()
}

func (service *productService) FindProductByCategoryId(ctx context.Context, id int) []response.ResponseProduct {
	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.productRepository.FindProductByCategoryId(ctx, tx, id)

	responseProducts := []response.ResponseProduct{}
	for _, v := range products {
		responseProducts = append(responseProducts, v.ToResponseProduct())
	}

	return responseProducts
}

func (service *productService) CreateProductByCategoryId(ctx context.Context, request request.RequestCreateProduct, id int) response.ResponseProduct {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product := domain.Product{}
	product.SetName(&request.Name)
	product.SetCategoryId(&request.CategoryId)
	product = service.productRepository.CreateProductByCategoryId(ctx, tx, product, id)

	return product.ToResponseProduct()
}

func (service *productService) UpdateProductByCategoryId(ctx context.Context, request request.RequestUpdateProduct) {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product := service.productRepository.FindById(ctx, tx, request.Id)
	product.SetName(&request.Name)

	service.productRepository.Update(ctx, tx, product)
}

func (service *productService) DeleteProductByCategoryId(ctx context.Context, request request.RequestDeleteProduct) {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product := service.productRepository.FindById(ctx, tx, request.Id)

	service.productRepository.Delete(ctx, tx, product)
}