package main

import (
	"github.com/Clementol/1gocrud/database"
	"github.com/Clementol/1gocrud/handlers"
	"github.com/Clementol/1gocrud/middlewares/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// func WithDb(fn func(c *gin.Context, db *mongo.Database)
//  db *mongo.Database func(c *gin.Context)  {
// 	return func(c *gin.Context) error {
// 		return fn(c, db)
// 	}
// }

func main() {
	godotenv.Load()
	router := gin.Default()
	router.Use(cors.CORSMiddleware())
	_, _ = database.ConnectToDatase()

	api := router.Group("/api/staff")
	{
		api.GET("/all", handlers.HandleGetAllEmployes)
		api.POST("/add", handlers.AddEmployeeHandler)
		api.PUT("/update/:id", handlers.UpdateEmployee)
		api.DELETE("/delete/:id", handlers.DeleteEmployee)
	}

	router.Run()
}
