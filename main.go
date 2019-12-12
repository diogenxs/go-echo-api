package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
	})

	e.GET("/ping", func(c echo.Context) error {
		type Message struct {
			Message string `json:"message"`
		}
		m := &Message{"pong"}
		return c.JSON(http.StatusOK, &m)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
