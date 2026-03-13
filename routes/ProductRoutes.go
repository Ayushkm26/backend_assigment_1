package routes

import (
	"Backend_assigment_1/controller"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(rg *gin.RouterGroup) {

	productController := controller.ProductController{}

	products := rg.Group("/products")
	{
		products.GET("/", productController.GetProducts)

	}
}
