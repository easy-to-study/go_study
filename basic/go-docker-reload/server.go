package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World From Echo!")
	})
	e.GET("/tom/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, "Hello,"+name+"!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
