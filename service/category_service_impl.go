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

type categoryService struct {
	db                 *sql.DB
	categoryRepository repository.CategoryRepository
	// nambahin validate
	validate *validator.Validate
}

func NewCategoryService(db *sql.DB, categoryRepository repository.CategoryRepository, validate *validator.Validate) *categoryService {
	return &categoryService{db, categoryRepository, validate}
}

func (service *categoryService) FindAll(ctx context.Context) []response.ResponseCategory {
	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.categoryRepository.FindAll(ctx, tx)

	responseCategories := []response.ResponseCategory{}
	for _, v := range categories {
		responseCategories = append(responseCategories, v.ToResponseCategory())
	}

	return responseCategories
}

func (service *categoryService) Create(ctx context.Context, request request.RequestCreateCategory) response.ResponseCategory {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category := domain.Category{}
	category.SetName(&request.Name)
	category = service.categoryRepository.Create(ctx, tx, category)

	return category.ToResponseCategory()
}

func (service *categoryService) Update(ctx context.Context, request request.RequestUpdateCategory) {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category := service.categoryRepository.FindById(ctx, tx, request.Id)
	category.SetName(&request.Name)

	service.categoryRepository.Update(ctx, tx, category)
}

func (service *categoryService) Delete(ctx context.Context, request request.RequestDeleteCategory) {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category := service.categoryRepository.FindById(ctx, tx, request.Id)

	service.categoryRepository.Delete(ctx, tx, category)
}

func (service *categoryService) FindById(ctx context.Context, id int) response.ResponseCategory {
	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category := service.categoryRepository.FindById(ctx, tx, id)

	return category.ToResponseCategory()
}

