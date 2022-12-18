package routes

import (
	"echo-curd/controller"

	"github.com/labstack/echo/v4"
)

var AllRoutes = func(e *echo.Echo) {
	e.GET("/", controller.List)
	e.POST("/create", controller.CreateTodo)
	e.GET("/todo/:id", controller.GetTodoById)
	e.POST("/todo/:id", controller.UpdateTodo)
	e.DELETE("/todo/:id", controller.DeleteTodoById)
}
