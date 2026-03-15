package services

import (
	"Backend_assigment_1/DatatbaseConnection"
	"Backend_assigment_1/kafka"
	"Backend_assigment_1/models"
	"errors"
	"time"
)

type Order struct {
}

func (o *Order) CreateOrders(input models.ReqOrder) (models.Order, []models.OrderItem, error) {
	db := DatatbaseConnection.DB

	if len(input.Items) == 0 {
		return models.Order{}, nil, errors.New("order must have at least one item")
	}
	tx := db.Begin()
	if tx.Error != nil {
		return models.Order{}, nil, tx.Error
	}

	var total float64
	var items []models.OrderItem
	for _, it := range input.Items {
		var product models.Product
		if err := tx.First(&product, it.ProductID).Error; err != nil {
			tx.Rollback()
			return models.Order{}, nil, errors.New("product not found")
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
		Status:      "pending",
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

	tx.Commit()
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

	err := kafka.PublishOrderCreated(event)
	if err != nil {
		return models.Order{}, nil, err
	}
	return order, items, nil
}
