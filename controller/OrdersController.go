package controller

import (
	"Backend_assigment_1/models"
	"net/http"
	"strconv"

	"Backend_assigment_1/services"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService *services.Order
}

func (o *OrderController) CreateOrders(c *gin.Context) {
	var input models.ReqOrder

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	order, items, err := o.OrderService.CreateOrders(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order": order,
		"items": items,
	})

}
func (o *OrderController) GetOrdersById(c *gin.Context) {
	// Get path param "id"
	idParam := c.Param("id")
	userID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call service layer
	orders, err := o.OrderService.GetOrdersByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	// If no orders found
	if len(orders) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No orders found for this user"})
		return
	}

	// Return orders
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
