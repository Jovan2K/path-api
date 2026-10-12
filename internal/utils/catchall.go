package utils

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func CatchAll(e *echo.Echo) {
	e.GET("/", func(c *echo.Context) error {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	})
	e.GET("/*", func(c *echo.Context) error {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	})
}
