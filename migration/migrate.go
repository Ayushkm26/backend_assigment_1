package migrations

import (
	"Backend_assigment_1/DatatbaseConnection"
	"Backend_assigment_1/models"
	"log"
)

func Migrate() {
	db := DatatbaseConnection.DB
	if err := db.AutoMigrate(&models.Product{}, &models.Order{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed")
}
