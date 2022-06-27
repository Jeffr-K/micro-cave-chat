package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func WebSocketRouter(c echo.Context) (e error) {
	return c.JSON(http.StatusOK, "true")
}
