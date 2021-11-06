package models

import (
	"time"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	LastName   string    `bson:"lastName"`
	FirstName  string    `bson:"firstName"`
	Email      string    `bson:"email"`
	Position   string    `bson:"position"`
	Department string    `bson:"department"`
	CreatedAt  time.Time `bson:"createdAt"`
	UpdatedAt  time.Time `bson:"updatedAt"`
}
