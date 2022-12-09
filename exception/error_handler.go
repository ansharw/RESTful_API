package exception

import (
	"belajar-rest-api/model/api"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	internalServerError(err, c)
}

func internalServerError(err error, c echo.Context) {
	c.JSON(http.StatusInternalServerError, api.ApiResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err.Error(),
	})
}

// without validator
// func PanicMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		defer func(c echo.Context) error {
// 			err := recover()

// 			notFoundError, ok := err.(NotFoundError)

// 			if ok {
// 				return c.JSON(http.StatusNotFound, api.ApiResponse{
// 					Code:   http.StatusNotFound,
// 					Status: "NOT FOUND",
// 					Data:   notFoundError.Error.Error(),
// 				})
// 			}
// 			internalServerError, ok := err.(InternalServerError)

// 			if ok {
// 				return c.JSON(http.StatusInternalServerError, api.ApiResponse{
// 					Code:   http.StatusInternalServerError,
// 					Status: "INTERNAL SERVER ERROR",
// 					Data:   internalServerError.Error.Error(),
// 				})
// 			}
// 			return c.NoContent(http.StatusInternalServerError)
// 		}(c)
// 		return next(c)
// 	}
// }

// WITH VALIDATOR
func PanicMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func(c echo.Context) error {
			err := recover()

			validationErrors, ok := err.(validator.ValidationErrors)
			fmt.Println("string 1", validationErrors)
			fmt.Println(ok)

			if ok {
				return c.JSON(http.StatusBadRequest, api.ApiResponse{
					Code:   http.StatusBadRequest,
					Status: "BAD REQUEST",
					Data:   validationErrors.Error(),
				})
			}

			notFoundError, ok := err.(NotFoundError)
			fmt.Println("string 2", notFoundError)
			fmt.Println(ok)

			if ok {
				return c.JSON(http.StatusNotFound, api.ApiResponse{
					Code:   http.StatusNotFound,
					Status: "NOT FOUND",
					Data:   notFoundError.Error.Error(),
				})
			}

			internalServerError, ok := err.(InternalServerError)
			fmt.Println("string 3", internalServerError)
			fmt.Println(ok)

			if ok {
				return c.JSON(http.StatusInternalServerError, api.ApiResponse{
					Code:   http.StatusInternalServerError,
					Status: "INTERNAL SERVER ERROR",
					Data:   internalServerError.Error.Error(),
				})
			}
			return c.NoContent(http.StatusInternalServerError)
		}(c)
		return next(c)
	}
}
