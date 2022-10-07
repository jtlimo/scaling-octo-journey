package adapters

import (
	"order/pkg/order/domain"
)

func Adapt(o *domain.Order) OrderResponse {
	return OrderResponse{
		Id:            o.Id,
		PaymentMethod: o.PaymentMethod,
		Status:        Kind(o.Status.Kind),
		Address:       o.Address,
		Itens:         toItens2(o.Itens),
		Merchant:      Merchant(o.Merchant),
	}
}

func AdaptToDomain(companyID string, orderRequestBody OrderRequestBody) *domain.Order {
	return &domain.Order{
		PaymentMethod: orderRequestBody.PaymentMethod,
		Merchant: domain.Merchant{
			Id: companyID,
		},
		Itens:   toItens(orderRequestBody.Itens),
		Address: orderRequestBody.Address,
	}
}

// FIXME: Learn how to implement this function using generics
func toItens2(itens []domain.Item) []Item {
	it := make([]Item, 0)
	for _, i := range itens {
		it = append(it, Item{
			Quantity: i.Quantity,
			Product:  ToProduct2(i.Product),
			Comment:  i.Comment,
		})
	}
	return it
}

func toItens(itens []Item) []domain.Item {
	it := make([]domain.Item, 0)
	for _, i := range itens {
		it = append(it, domain.Item{
			Quantity: i.Quantity,
			Product:  ToProduct(i.Product),
			Comment:  i.Comment,
		})
	}
	return it
}

// FIXME: Learn how to implement this function using generics
func ToProduct2(product domain.Product) Product {
	return Product{
		Id: product.Id,
	}
}

func ToProduct(product Product) domain.Product {
	return domain.Product{
		Id: product.Id,
	}
}
