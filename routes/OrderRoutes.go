package routes

import (
	"Backend_assigment_1/controller"
	"Backend_assigment_1/services"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(rg *gin.RouterGroup) {

	orderController := controller.OrderController{
		OrderService: &services.Order{}, // doinng  dependecies injection creating controller service and injecting orderService in to it as it is depending upon it
	}

	orders := rg.Group("/orders")
	{
		orders.POST("/createorder", orderController.CreateOrders)
		orders.GET("/getOrdersById/:id", orderController.GetOrdersById)

	}

}
