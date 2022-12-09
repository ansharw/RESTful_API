package service

import (
	"belajar-rest-api/model/request"
	"belajar-rest-api/model/response"
	"context"
)

type CategoryService interface {
	FindAll(ctx context.Context) []response.ResponseCategory
	Create(ctx context.Context, request request.RequestCreateCategory) response.ResponseCategory
	Update(ctx context.Context, request request.RequestUpdateCategory)
	Delete(ctx context.Context, request request.RequestDeleteCategory)

	FindById(ctx context.Context, id int) response.ResponseCategory
}
