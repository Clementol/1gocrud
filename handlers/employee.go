package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Clementol/1gocrud/controllers"
	"github.com/Clementol/1gocrud/database"
	"github.com/Clementol/1gocrud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

func HandleGetAllEmployes(c *gin.Context) {
	employees, err := controllers.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"staffs": employees})

}

func AddEmployeeHandler(c *gin.Context) {

	employeeCollection := database.EmployeeCollection()
	var employee models.Employee
	var err error

	if err := c.ShouldBind(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var staffEmailExist bson.M
	_ = employeeCollection.FindOne(context.TODO(),
		bson.M{"email": employee.Email}).Decode(&staffEmailExist)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err.Error()})
	// 	return
	// }

	if staffEmailExist != nil {

		message := "staff with email already exists! "
		c.JSON(http.StatusConflict, gin.H{"error": message})
		return
	}
	data := bson.M{
		"lastName":   employee.LastName,
		"firstName":  employee.FirstName,
		"email":      employee.Email,
		"position":   employee.Position,
		"department": employee.Department,
		"createdAt":  time.Now(),
		"updatedAt":  time.Now(),
	}

	_, err = employeeCollection.InsertOne(context.TODO(), data)
	if err != nil {
		message := `staff not added ` + err.Error()
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusCreated, "staff added")

}

func UpdateEmployee(c *gin.Context) {

	employeeCollection := database.EmployeeCollection()
	var employee models.Employee
	var err error
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	err = c.ShouldBind(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEmployee := bson.M{}

	data := bson.M{
		"lastName":   employee.LastName,
		"firstName":  employee.FirstName,
		"email":      employee.Email,
		"position":   employee.Position,
		"department": employee.Department,
		"updatedAt":  time.Now(),
	}
	filter := bson.M{"_id": bson.M{"$eq": id}, "email": employee.Email}
	update := bson.M{"$set": data}

	err = employeeCollection.FindOneAndUpdate(context.TODO(),
		filter, update).Decode(&updatedEmployee)

	if err != nil {
		message := `staff with email does not exist!` + err.Error()
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"updatedStaff": updatedEmployee})

}

func DeleteEmployee(c *gin.Context) {

	employeeCollection := database.EmployeeCollection()

	var err error
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	deletedEmployee := bson.M{}

	filter := bson.M{"_id": bson.M{"$eq": id}}

	err = employeeCollection.FindOneAndDelete(context.TODO(),
		filter).Decode(&deletedEmployee)

	if err != nil {
		message := `staff with email does not exist! ` + err.Error()
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusOK, "staff deleted")

}
