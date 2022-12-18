package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"echo-curd/models"
)

var NewTodo models.Todo

func List(c echo.Context) error {
	NewTodo := models.GetAllTodos()
	return c.JSON(http.StatusOK, NewTodo)
}

func CreateTodo(c echo.Context) error {
	newTodo := new(models.Todo)
	if err := c.Bind(newTodo); err != nil {
		return err
	}
	b := newTodo.CreateTodo()
	return c.JSON(http.StatusCreated, b)
}

func GetTodoById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 0, 0)
	if err != nil {

		return c.String(http.StatusOK, "Unable to parse.. Make sure the id is int")
	}
	todo, _ := models.GetTodoById(id)
	// fmt.Println(todo)
	if todo != nil {
		return c.JSON(http.StatusOK, todo)
	}
	return c.String(http.StatusMovedPermanently, "Unable to fetech the details")

}

func DeleteTodoById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 0, 0)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to parse.. Make sure the id is int")
	}
	models.DeleteTodo(id)

	return c.String(http.StatusOK, "Deleted Successfully")
}

func UpdateTodo(c echo.Context) error {
	todo := new(models.Todo)
	if err := c.Bind(&todo); err != nil {
		return err
	}

	oldTodo, db := models.GetTodoById(int64(todo.Id))

	// oldTodo.Name = todo.Name
	// oldTodo.Description = todo.Description
	oldTodo = todo
	db.Save(&oldTodo)
	return c.JSON(http.StatusOK, oldTodo)
}
