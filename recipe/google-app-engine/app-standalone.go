// +build !appengine,!appenginevm

package main

import (
	"github.com/Hunter-Dolan/echo"
	"github.com/Hunter-Dolan/echo/engine/standard"
	"github.com/Hunter-Dolan/echo/middleware"
)

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Use(middleware.Static("public"))

	return e
}

func main() {
	e.Run(standard.New(":8080"))
}
