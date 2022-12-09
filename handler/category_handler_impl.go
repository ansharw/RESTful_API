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

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *categoryHandler {
	return &categoryHandler{categoryService: categoryService}
}

func (handler *categoryHandler) FindAll(c echo.Context) error {
	responseCategories := handler.categoryService.FindAll(c.Request().Context())

	// return c.JSON(http.StatusOK,
	// 	map[string]interface{}{
	// 		"code":   200,
	// 		"status": "OK",
	// 		"data":   responseCategories,
	// 	},
	// )

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseCategories,
	})
}

func (handler *categoryHandler) Create(c echo.Context) error {
	requestCategory := new(request.RequestCreateCategory)
	c.Bind(requestCategory)

	responseCategory := handler.categoryService.Create(c.Request().Context(), *requestCategory)

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseCategory,
	})
}

func (handler *categoryHandler) Update(c echo.Context) error {
	requestCategory := new(request.RequestUpdateCategory)
	err := c.Bind(requestCategory)
	helper.PanicIfError(err)

	handler.categoryService.Update(c.Request().Context(), *requestCategory)

	return c.JSON(http.StatusOK, api.ApiResponse{		
		Code:   200,
		Status: "OK",
		Data:   requestCategory,
	})
}

func (handler *categoryHandler) Delete(c echo.Context) error {
	requestCategory := new(request.RequestDeleteCategory)
	err := c.Bind(requestCategory)
	helper.PanicIfError(err)

	handler.categoryService.Delete(c.Request().Context(), *requestCategory)

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
	})
}

func (handler *categoryHandler) FindById(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	helper.PanicIfError(err)

	responseCategory := handler.categoryService.FindById(c.Request().Context(), int(id))

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data: responseCategory,
	})
}
