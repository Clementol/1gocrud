package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseInterface interface {
	EmployeeCollection() *mongo.Collection
}

var db *mongo.Database

func ConnectToDatase() {
	mongodbUri := os.Getenv("MONGO_URI")

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI(mongodbUri))
	if err != nil {

		panic(err.Error())
	}

	// defer client.Disconnect(ctx)
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Printf(`Not connected to database! -> %v`, err.Error())
		// panic(err.Error())

	} else {
		log.Printf(`Connected to database!`)
	}

	db = client.Database("company_renaissance")
	// defer cancel()
	// return db, ctx

}

func EmployeeCollection() *mongo.Collection {
	return db.Collection("employees")
}
