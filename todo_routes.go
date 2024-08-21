package routes

import (
	"todo-service/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.Engine, todoCollection *mongo.Collection) {
	router.POST("/todo/v1/entry", func(c *gin.Context) {
		controllers.CreateTodoEntry(c, todoCollection)
	})
	router.GET("/todo/v1/", func(c *gin.Context) {
		controllers.GetTodos(c, todoCollection)
	})
	router.PUT("/todo/v1/entry/:id", func(c *gin.Context) {
		controllers.UpdateTodoEntry(c, todoCollection)
	})
	router.DELETE("/todo/v1/entry/:id", func(c *gin.Context) {
		controllers.DeleteTodoEntry(c, todoCollection)
	})
}
