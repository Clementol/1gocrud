package main

import (
	"github.com/Clementol/1gocrud/database"
	"github.com/Clementol/1gocrud/handlers"
	"github.com/Clementol/1gocrud/middlewares/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := gin.Default()
	router.Use(cors.CORSMiddleware())
	api := router.Group("/api/staff")
	{
		api.GET("/all", handlers.HandleGetAllEmployes)
		api.POST("/add", handlers.AddEmployeeHandler)
		api.PUT("/update/:id", handlers.UpdateEmployee)
		api.DELETE("/delete/:id", handlers.DeleteEmployee)
	}
	database.ConnectToDatase()

	router.Run()
}
