package routes

import (
	"todo-service/depedency"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.Engine, controller depedency.TodoController) {
	router.POST("/todo/v1/entry", func(c *gin.Context) {
		controller.CreateTodoEntry(c)
	})
	// router.GET("/todo/v1/", func(c *gin.Context) {
	// 	controllers.GetTodos(c, todoCollection)
	// })
	// router.PUT("/todo/v1/entry/:id", func(c *gin.Context) {
	// 	controllers.UpdateTodoEntry(c, todoCollection)
	// })
	// router.DELETE("/todo/v1/entry/:id", func(c *gin.Context) {
	// 	controllers.DeleteTodoEntry(c, todoCollection)
	// })
}
