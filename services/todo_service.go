package services

import (
	"context"
	"time"
	"todo-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoService struct {
	TodoCollection *mongo.Collection
}

var ctx = context.Background()

func (t *TodoService) CreateTodo(todo models.Todo) error {
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	result, err := t.TodoCollection.InsertOne(ctx, todo)
	if result.InsertedID == nil && err != nil {
		return err
	}
	todo.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func GetTodos(status string, todoCollection *mongo.Collection) ([]models.Todo, error) {
	var todos []models.Todo
	filter := bson.M{"deletedAt": bson.M{"$exists": false}}
	if status != "" {
		filter["status"] = status
	}
	cursor, err := todoCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var todo models.Todo
		err := cursor.Decode(&todo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func UpdateTodo(id string, updatedTodo models.Todo, todoCollection *mongo.Collection) (*mongo.UpdateResult, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	updatedTodo.UpdatedAt = time.Now()
	update := bson.M{"$set": updatedTodo}
	result, err := todoCollection.UpdateOne(ctx, filter, update)
	return result, err
}

func DeleteTodo(id string, todoCollection *mongo.Collection) (*mongo.DeleteResult, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	result, err := todoCollection.DeleteOne(ctx, filter)
	return result, err
}
