package main

import (
	"log"
	"todo-service/database"
	"todo-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB cloud
	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer database.DisconnectDB()

	// Get the "todos" collection
	todoCollection := database.GetCollection("todos")

	// Initialize Gin engine
	router := gin.Default()

	// Inject todoCollection to handlers
	routes.TodoRoutes(router, todoCollection)

	// Run the server on port 8080
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
