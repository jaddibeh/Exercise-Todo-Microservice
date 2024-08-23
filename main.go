package main

import (
	"log"
	"todo-service/controllers"
	"todo-service/database"
	"todo-service/routes"
	"todo-service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB cloud
	client, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer database.DisconnectDB()

	// Get the "todos" collection
	todoCollection := client.Database("todo").Collection("todos")
	todoService := services.TodoService{TodoCollection: todoCollection}
	controller := controllers.TodoController{TodoService: &todoService}

	// Initialize Gin engine
	router := gin.Default()

	// Inject todoCollection to handlers
	routes.TodoRoutes(router, &controller)

	// Run the server on port 8080
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
