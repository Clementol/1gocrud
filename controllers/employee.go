package controllers

import (
	"context"
	"log"

	"github.com/Clementol/1gocrud/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllEmployees() []*primitive.M {

	var employees []*bson.M
	col := database.ConnectToDatase()

	// defer cancel()
	// defer client.Disconnect(ctx)

	employeesCollection := col.Collection("employees")
	cur, err := employeesCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cur.All(context.TODO(), &employees); err != nil {
		log.Fatal(err.Error())
	}
	// defer cur.Close()
	// for cur.Next(context.TODO()) {
	// 	var elem models.Employee
	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// 	employees = append(employees, &elem)
	// }

	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err.Error())
	// }
	// cur.Close(context.TODO())
	return employees

}
