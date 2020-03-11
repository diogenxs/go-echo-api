package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "public",
		HTML5: true,
	}))

	api := e.Group("/api")
	v1 := api.Group("/v1")
	// Middleware
	v1.Use(middleware.RequestID())
	v1.Use(middleware.Logger())
	v1.Use(middleware.Recover())

	v1.GET("/ping", Ping)

	return e
}
