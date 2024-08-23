package depedency

import (
	"todo-service/models"

	"github.com/gin-gonic/gin"
)

type TodoService interface {
	CreateTodo(todo models.Todo) error
}

type TodoController interface {
	CreateTodoEntry(c *gin.Context) error
}
