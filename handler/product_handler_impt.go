package handler

import (
	"belajar-rest-api/helper"
	"belajar-rest-api/model/api"
	"belajar-rest-api/model/request"
	"belajar-rest-api/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{productService: productService}
}

func (handler *productHandler) FindAll(c echo.Context) error {
	responseProducts := handler.productService.FindAll(c.Request().Context())

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseProducts,
	})
}

func (handler *productHandler) Create(c echo.Context) error {
	requestProduct := new(request.RequestCreateProduct)
	c.Bind(requestProduct)

	responseProduct := handler.productService.Create(c.Request().Context(), *requestProduct)

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseProduct,
	})
}

func (handler *productHandler) Update(c echo.Context) error {
	requestProduct := new(request.RequestUpdateProduct)
	err := c.Bind(requestProduct)
	helper.PanicIfError(err)

	handler.productService.Update(c.Request().Context(), *requestProduct)

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   requestProduct,
	})
}

func (handler *productHandler) Delete(c echo.Context) error {
	requestProduct := new(request.RequestDeleteProduct)
	err := c.Bind(requestProduct)
	helper.PanicIfError(err)

	handler.productService.Delete(c.Request().Context(), *requestProduct)

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
	})
}

func (handler *productHandler) FindById(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	helper.PanicIfError(err)

	responseProduct := handler.productService.FindById(c.Request().Context(), int(id))

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseProduct,
	})
}

func (handler *productHandler) FindProductByCategoryId(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	helper.PanicIfError(err)

	responseProducts := handler.productService.FindProductByCategoryId(c.Request().Context(), int(id))

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseProducts,
	})
}
