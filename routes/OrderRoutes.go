package routes

import (
	"Backend_assigment_1/controller"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(rg *gin.RouterGroup) {

	orderController := controller.OrderController{}

	orders := rg.Group("/orders")
	{
		orders.GET("/createOrder", orderController.CreateOrders)

	}
}
