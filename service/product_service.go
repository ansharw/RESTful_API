package service

import (
	"belajar-rest-api/model/request"
	"belajar-rest-api/model/response"
	"context"
)

type ProductService interface {
	FindAll(ctx context.Context) []response.ResponseProduct
	Create(ctx context.Context, request request.RequestCreateProduct) response.ResponseProduct
	Update(ctx context.Context, request request.RequestUpdateProduct)
	Delete(ctx context.Context, request request.RequestDeleteProduct)

	FindById(ctx context.Context, id int) response.ResponseProduct
	FindProductByCategoryId(ctx context.Context, id int) []response.ResponseProduct
}
