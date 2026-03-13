package services

import "Backend_assigment_1/models"

type Order struct {
}

func (o *Order) CreateOrders(input models.ReqOrder) (models.ReqOrder, error) {
	return models.ReqOrder{}, nil
}
