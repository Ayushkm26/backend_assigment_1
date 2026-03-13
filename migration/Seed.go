package migrations

import (
	"Backend_assigment_1/DatatbaseConnection"
	"Backend_assigment_1/models"
	"log"
	"time"
)

func SeedProducts() {
	db := DatatbaseConnection.DB

	// Check if products already exist
	var count int64
	db.Model(&models.Product{}).Count(&count)
	if count > 0 {
		log.Println("Products already seeded, skipping...")
		return
	}

	// Create 10 sample products
	products := []models.Product{
		{
			Name:        "Wireless Mouse",
			Price:       12.99,
			Stock:       50,
			Description: "Ergonomic wireless mouse with adjustable DPI.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Mechanical Keyboard",
			Price:       45.50,
			Stock:       30,
			Description: "RGB backlit mechanical keyboard with blue switches.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "USB-C Hub",
			Price:       25.00,
			Stock:       40,
			Description: "7-in-1 USB-C hub with HDMI, USB, SD card ports.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Gaming Headset",
			Price:       39.99,
			Stock:       20,
			Description: "Over-ear gaming headset with noise cancellation.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Laptop Stand",
			Price:       18.75,
			Stock:       60,
			Description: "Aluminum adjustable laptop stand for desk use.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "External Hard Drive",
			Price:       75.00,
			Stock:       15,
			Description: "1TB external hard drive with USB 3.0.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Webcam HD",
			Price:       30.00,
			Stock:       25,
			Description: "1080p HD webcam with built-in microphone.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Portable Charger",
			Price:       22.50,
			Stock:       70,
			Description: "10000mAh portable charger for phones and tablets.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Bluetooth Speaker",
			Price:       35.00,
			Stock:       45,
			Description: "Portable Bluetooth speaker with deep bass.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Smartwatch",
			Price:       99.99,
			Stock:       10,
			Description: "Smartwatch with heart rate monitor and notifications.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	// Insert into database
	if err := db.Create(&products).Error; err != nil {
		log.Fatal("Failed to seed products:", err)
	}

	log.Println("Seeded 10 products successfully!")
}
