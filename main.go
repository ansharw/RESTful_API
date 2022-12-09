package main

import (
	"belajar-rest-api/database"
	"belajar-rest-api/exception"
	"belajar-rest-api/handler"
	// "belajar-rest-api/middleware"
	"belajar-rest-api/repository"
	"belajar-rest-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	db := database.GetConnection()
	validator := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(db, categoryRepository, validator)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	echo := echo.New()
	api := echo.Group("/api")
	// echo.Use(exception.PanicMiddleWare, middleware.AuthMiddleware)
	echo.Use(exception.PanicMiddleWare)
	api.GET("/categories", categoryHandler.FindAll)
	api.POST("/categories", categoryHandler.Create)
	api.PUT("/categories", categoryHandler.Update)
	api.DELETE("/categories", categoryHandler.Delete)
	api.GET("/categories/:id", categoryHandler.FindById)

	echo.Logger.Fatal(echo.Start(":3000"))
}
