package controllers

import (
	"context"

	"github.com/Clementol/1gocrud/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllEmployees() ([]*primitive.M, error) {

	var employees []*bson.M
	var err error
	col, ctx := database.ConnectToDatase()
	employeesCollection := col.Collection("employees")
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	findOption := []bson.M{
		{"$project": bson.M{
			"id": "$_id",
			"name": bson.M{
				"$concat": []string{
					"$lastName", " ", "$firstName",
				},
			},
			"lastName":   "$lastName",
			"firstName":  "$firstName",
			"email":      "$email",
			"position":   "$position",
			"department": "$department",
		}}}

	// employeesCollection := col.Collection("employees")
	cur, err := employeesCollection.Aggregate(ctx, findOption)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}

	if err = cur.All(context.TODO(), &employees); err != nil {
		// log.Fatal(err.Error())
		return nil, err
	}

	return employees, err

}
