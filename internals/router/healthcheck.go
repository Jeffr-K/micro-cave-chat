package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheckRouter(c echo.Context) (e error) {
	return c.JSON(http.StatusOK, "true")
}
