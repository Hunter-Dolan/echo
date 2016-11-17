package main

import (
	"net/http"

	"github.com/Hunter-Dolan/echo"
	"github.com/Hunter-Dolan/echo/engine/standard"
	"github.com/Hunter-Dolan/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	// Start server
	e.Run(standard.New(":1323"))
}
