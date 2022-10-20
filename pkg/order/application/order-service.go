package application

import (
	"order/pkg/order/domain"
	"order/pkg/order/repository"
)

type Application struct {
	repository *repository.OrderRepository
}

func New(repository *repository.OrderRepository) *Application {
	return &Application{
		repository: repository,
	}
}

func (a *Application) CreateOrder(order *domain.Order) error {
	err := a.repository.Insert(order)
	if err != nil {
		return err
	}
	return nil
}
