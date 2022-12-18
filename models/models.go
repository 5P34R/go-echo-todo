package models

import (
	"echo-curd/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Todo struct {
	gorm.Model
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}

func init() {
	config.Connect()
	db = config.GetTodo()
	db.AutoMigrate(&Todo{})
}

func GetAllTodos() []Todo {
	var todos []Todo
	db.Find(&todos)
	return todos
}

func (todo *Todo) CreateTodo() *Todo {
	db.Create(&todo)
	return todo
}

func GetTodoById(id int64) (*Todo, *gorm.DB) {
	var todos Todo
	var count int64
	db := db.Where("ID=?", id).Find(&todos).Count(&count)
	// fmt.Println(count)
	if count == 0 {
		return nil, db
	}
	return &todos, db
}

func DeleteTodo(id int64) Todo {
	var todo Todo
	db.Where("ID=?", id).Delete(&todo)
	return todo
}
