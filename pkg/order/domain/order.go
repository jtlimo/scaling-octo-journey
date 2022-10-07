package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Kind string

const (
	Pending   Kind = "pending"
	Preparing Kind = "preparing"
	Canceled  Kind = "canceled"
	Delivered Kind = "delivered"
)

type Order struct {
	Id            string
	PaymentMethod string
	Status        Status
	Address       string
	Itens         []Item
	Merchant      Merchant
}

type Status struct {
	Kind   Kind
	Reason string
}

type Item struct {
	Quantity int
	Product  Product
	Comment  string
}

type Product struct {
	Id string
}

type Merchant struct {
	Id string
}

var GenerateNewUUID = uuid.NewString

func (o *Order) cancel(reason string) error {
	if o.Status.Kind != Delivered {
		o.Status.Kind = Canceled
		o.Status.Reason = reason
		return nil
	}
	return errors.New("cannot cancel the order")
}

func (o *Order) delivered() {
	o.Status.Kind = Delivered
}

// TODO: create a method to prepare a order - put an order to preparing status
func (o *Order) prepare() {

}

func New(paymentMethod string, address string, itens []Item, merchant Merchant) (*Order, error) {
	duuid := GenerateNewUUID()
	order := &Order{
		Id:            duuid,
		PaymentMethod: paymentMethod,
		Status:        Status{Kind: Pending, Reason: ""},
		Address:       address,
		Itens:         itens,
		Merchant:      merchant,
	}

	return order, nil
}
