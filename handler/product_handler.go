package handler

import "github.com/labstack/echo/v4"

type ProductHandler interface {
	FindAll(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error

	FindById(c echo.Context) error

	FindProductByCategoryId(c echo.Context) error
	CreateProductByCategoryId(c echo.Context) error
	UpdateProductByCategoryId(c echo.Context) error
	DeleteProductByCategoryId(c echo.Context) error
}
