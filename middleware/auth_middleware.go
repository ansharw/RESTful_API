package middleware

import (
	"belajar-rest-api/model/api"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("X-API-Key") == "Secret" {
			return next(c)
		} else {
			return c.JSON(http.StatusUnauthorized, api.ApiResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUHTORIZED",
			})
		}
	}
}