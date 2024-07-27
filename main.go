package main

import (
	"GoAPI/config"
	"GoAPI/controllers"
	"GoAPI/models"
	"GoAPI/repositories"
	"GoAPI/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	_ "GoAPI/docs"
)

// @title Employee API
// @version 1.0
// @description This is a sample Employee CRUD API.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Connect to database
	db, err := gorm.Open(sqlserver.Open(cfg.ConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Ensure the table structure matches our model
	db.Set("gorm:table_options", "").AutoMigrate(&models.Employee{})

	// Set up repository, service, and controller
	repo := repositories.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repo)
	controller := controllers.NewEmployeeController(service)

	// Set up Gin router
	r := gin.Default()

	// API routes
	v1 := r.Group("/api/v1")
	{
		employees := v1.Group("/employees")
		{
			employees.POST("", controller.CreateEmployee)
			employees.GET("", controller.GetAllEmployees)
			employees.GET("/:id", controller.GetEmployee)
			employees.PUT("/:id", controller.UpdateEmployee)
			employees.DELETE("/:id", controller.DeleteEmployee)
		}
	}

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the server
	r.Run(":8080")
}
