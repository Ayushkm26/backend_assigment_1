package controller

import (
	"Backend_assigment_1/models"
	"net/http"

	"Backend_assigment_1/services"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService services.Order
}

func (o *OrderController) CreateOrders(c *gin.Context) {
	var input models.ReqOrder

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	products, err := o.OrderService.CreateOrders(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}
