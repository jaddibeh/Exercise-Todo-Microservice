package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
    ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Title       string             `json:"title,omitempty" bson:"title,omitempty"`
    Description string             `json:"description,omitempty" bson:"description,omitempty"`
    Status      string             `json:"status,omitempty" bson:"status,omitempty"`
    CreatedAt   time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
    UpdatedAt   time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
    DeletedAt   *time.Time         `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}
