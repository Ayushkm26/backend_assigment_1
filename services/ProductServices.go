package services

import (
	"Backend_assigment_1/DatatbaseConnection"
	"Backend_assigment_1/models"
)

type Product struct{}

func (p *Product) GetProducts() ([]models.RespProduct, error) {

	var products []models.RespProduct

	query := `
		SELECT id, name, price, description,stock
		FROM products where stock>0
	`

	if err := DatatbaseConnection.DB.Raw(query).Scan(&products).Error; err != nil {
		return nil, err
	}

	return products, nil

}
