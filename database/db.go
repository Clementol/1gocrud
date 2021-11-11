package database

import (
	"context"
	"log"
	"os"
	"time"

	// "time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// type DatabaseInterface interface {
// 	ConnectToDatase()
// }

func ConnectToDatase() (*mongo.Database, context.Context) {
	mongodbUri := os.Getenv("MONGO_URI")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(mongodbUri))
	if err != nil {
		panic(err.Error())
	}

	// defer client.Disconnect(ctx)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Printf(`Not connected to database! -> %v`, err.Error())
		// panic(err.Error())

	} else {
		log.Printf(`Connected to database!`)
	}

	col := client.Database("company_renaissance")
	// defer cancel()
	return col, ctx

}
