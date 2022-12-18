package main

import (
	"echo-curd/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	routes.AllRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
