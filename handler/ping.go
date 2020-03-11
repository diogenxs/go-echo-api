package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	response := map[string]string{"message": "pong"}
	return c.JSON(http.StatusOK, response)
}
