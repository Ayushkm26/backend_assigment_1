package controller

import (
	"net/http"

	"Backend_assigment_1/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService services.Product
}

func (pc *ProductController) GetProducts(c *gin.Context) {

	products, err := pc.ProductService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}
