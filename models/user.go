package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User data type
type User struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"created_at"`
	UpdateAt  time.Time          `json:"update_at"`
}

// Users list
type Users []*User
