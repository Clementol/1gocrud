package models

import (
	"time"
)

type Employee struct {
	LastName   string    `bson:"lastName" validate:"required,min=2,max=20"`
	FirstName  string    `bson:"firstName" validate:"required,min=2,max=20"`
	Email      string    `bson:"email" validate:"required,email"`
	Position   string    `bson:"position" validate:"required,min=2,max=50"`
	Department string    `bson:"department" validate:"required,min=2,max=50"`
	CreatedAt  time.Time `bson:"createdAt"`
	UpdatedAt  time.Time `bson:"updatedAt"`
}
