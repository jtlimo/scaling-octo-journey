package repository

import (
	"errors"
	"order/pkg/order/domain"
)

type OrderRepository struct {
	orders []*domain.Order
}

func New() (repo *OrderRepository) {
	return &OrderRepository{
		orders: make([]*domain.Order, 0),
	}
}

func (or *OrderRepository) Insert(order *domain.Order) error {
	or.orders = append(or.orders, order)
	return nil
}

func (or *OrderRepository) getByID(id string) (*domain.Order, error) {
	for _, order := range or.orders {
		if order.Id == id {
			return order, nil
		}
	}
	return nil, errors.New("order not found")
}
