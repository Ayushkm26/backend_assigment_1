package services

import (
	"Backend_assigment_1/DatatbaseConnection"
	"Backend_assigment_1/kafka"
	"Backend_assigment_1/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order struct {
}

func (o *Order) CreateOrders(input models.ReqOrder) (models.Order, []models.OrderItem, error) {
	db := DatatbaseConnection.DB

	fmt.Println("=============================")
	fmt.Println("UserID received:", input.UserID)
	fmt.Println("Items received:", input.Items)
	fmt.Println("=============================")

	if input.UserID == 0 {
		return models.Order{}, nil, errors.New("invalid user ID")
	}

	if len(input.Items) == 0 {
		return models.Order{}, nil, errors.New("order must have at least one item")
	}

	tx := db.Begin()

	var user models.User
	tx.Where("id = ?", input.UserID).First(&user)

	fmt.Println("=============================")
	fmt.Println("User.ID after query:", user.ID)
	fmt.Println("User.Name:", user.Name)
	fmt.Println("User.Email:", user.Email)
	fmt.Println("=============================")

	if user.ID == 0 {
		tx.Rollback()
		fmt.Println("🔴 USER NOT FOUND — returning error")
		return models.Order{}, nil, errors.New("user not found")
	}

	fmt.Println("✅ User found — continuing order creation")

	var total float64
	var items []models.OrderItem

	for _, it := range input.Items {
		if it.Qty <= 0 {
			tx.Rollback()
			return models.Order{}, nil, errors.New("item quantity must be greater than zero")
		}

		// ✅ Reliable product check — same pattern as user check
		var product models.Product
		tx.First(&product, it.ProductID)
		if product.ID == 0 {
			tx.Rollback()
			return models.Order{}, nil, fmt.Errorf("product with ID %d not found", it.ProductID)
		}

		if product.Stock < it.Qty {
			tx.Rollback()
			return models.Order{}, nil, fmt.Errorf("insufficient stock for product ID %d", it.ProductID)
		}

		total += product.Price * float64(it.Qty)
		items = append(items, models.OrderItem{
			ProductID: it.ProductID,
			Qty:       it.Qty,
			Price:     product.Price,
		})
	}

	order := models.Order{
		UserID:      input.UserID,
		Status:      "created",
		TotalAmount: total,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return models.Order{}, nil, err
	}

	for i := range items {
		items[i].OrderID = order.ID
	}

	if err := tx.Create(&items).Error; err != nil {
		tx.Rollback()
		return models.Order{}, nil, err
	}

	for _, it := range items {
		if err := tx.Model(&models.Product{}).
			Where("id = ?", it.ProductID).
			UpdateColumn("stock", gorm.Expr("stock - ?", it.Qty)).Error; err != nil {
			tx.Rollback()
			return models.Order{}, nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return models.Order{}, nil, err
	}

	var eventItems []models.OrderItemEvent
	for _, it := range items {
		eventItems = append(eventItems, models.OrderItemEvent{
			ProductID: it.ProductID,
			Qty:       it.Qty,
		})
	}

	event := models.OrderCreatedEvent{
		OrderID: order.ID,
		UserID:  order.UserID,
		Items:   eventItems,
	}

	if err := kafka.PublishOrderCreated(event); err != nil {
		return models.Order{}, nil, err
	}

	return order, items, nil
}

func (s *Order) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	db := DatatbaseConnection.DB
	var orders []models.Order

	if err := db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}
