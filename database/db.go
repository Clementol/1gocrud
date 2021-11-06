package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// type DatabaseInterface interface {
// 	ConnectToDatase()
// }

func ConnectToDatase() *mongo.Database {
	mongodbUri := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI(mongodbUri))
	if err != nil {
		panic(err.Error())
	}

	// var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	// defer client.Disconnect(ctx)
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Printf(`Not connected to database! -> %v`, err.Error())
		// panic(err.Error())

	} else {
		log.Printf(`Connected to database!`)
	}

	col := client.Database("company_renaissance")
	// defer cancel()
	return col

}
