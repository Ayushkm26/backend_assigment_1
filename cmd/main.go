package main

import (
	"Backend_assigment_1/DatatbaseConnection"
	"Backend_assigment_1/kafka"
	migrations "Backend_assigment_1/migration"
	"Backend_assigment_1/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	err := DatatbaseConnection.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	migrations.Migrate()
	migrations.SeedProducts()

	kafka.InitProducer()

	router := gin.Default()

	// this allows the requets from the external  api of java to hit the golang api
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:9000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")

	routes.ProductRoutes(api)
	routes.OrderRoutes(api)

	log.Println("Database connected and routes registered successfully")

	router.Run(":8080")
}
