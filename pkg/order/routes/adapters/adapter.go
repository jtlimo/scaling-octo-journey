package adapters

import (
	"encoding/json"
	"order/pkg/order/domain"
)

func Adapt(o *domain.Order) OrderResponse {
	itens := toItens[domain.Item, Item](o.Itens)
	return OrderResponse{
		Id:            o.Id,
		PaymentMethod: o.PaymentMethod,
		Status:        Kind(o.Status.Kind),
		Address:       o.Address,
		Itens:         itens,
		Merchant:      Merchant(o.Merchant),
	}
}

func AdaptToDomain(companyID string, orderRequestBody OrderRequestBody) *domain.Order {
	itens := toItens[Item, domain.Item](orderRequestBody.Itens)
	return &domain.Order{
		PaymentMethod: orderRequestBody.PaymentMethod,
		Merchant: domain.Merchant{
			Id: companyID,
		},
		Itens:   itens,
		Address: orderRequestBody.Address,
	}
}

type Itens interface {
	Item | domain.Item
}

func toItens[IN Itens, OUT Itens](items []IN) []OUT {
	it := make([]OUT, 0)
	for _, i := range items {
		var m OUT
		mm, _ := json.Marshal(i)
		json.Unmarshal(mm, &m)
		it = append(it, m)
	}
	return it
}
