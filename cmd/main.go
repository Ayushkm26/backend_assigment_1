package main

import (
	"Backend_assigment_1/DatatbaseConnection"
	"Backend_assigment_1/migration"
	"Backend_assigment_1/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	err := DatatbaseConnection.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	migrations.Migrate()
	migrations.SeedProducts()
	router := gin.Default()
	api := router.Group("/api")

	routes.ProductRoutes(api)
	routes.OrderRoutes(api)
	log.Println("Database connected and routes registered successfully")
	router.Run(":8000")
}
