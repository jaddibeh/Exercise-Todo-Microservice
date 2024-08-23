package controllers

import (
	"net/http"
	"todo-service/depedency"
	"todo-service/models"
	"todo-service/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoController struct {
	TodoService depedency.TodoService
}

func (t *TodoController) CreateTodoEntry(c *gin.Context) error {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	err := t.TodoService.CreateTodo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	c.Status(http.StatusCreated)
	return nil
}

func GetTodos(c *gin.Context, todoCollection *mongo.Collection) {
	status := c.Query("status")
	todos, err := services.GetTodos(status, todoCollection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func UpdateTodoEntry(c *gin.Context, todoCollection *mongo.Collection) {
	id := c.Param("id")
	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := services.UpdateTodo(id, updatedTodo, todoCollection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteTodoEntry(c *gin.Context, todoCollection *mongo.Collection) {
	id := c.Param("id")
	result, err := services.DeleteTodo(id, todoCollection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
