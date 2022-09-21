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
	kind   Kind
	reason string
}

type Item struct {
	quantity int
	product  Product
	comment  string
}

type Product struct {
	id   string
	name string
}

type Merchant struct {
	id      string
	name    string
	address string
}

var GenerateNewUUID = uuid.NewString

func (o *Order) cancel(reason string) error {
	if o.Status.kind != Delivered {
		o.Status.kind = Canceled
		o.Status.reason = reason
		return nil
	}
	return errors.New("cannot cancel the order")
}

func (o *Order) delivered() {
	o.Status.kind = Delivered
}

// TODO: create a method to prepare a order - put an order to preparing status
func (o *Order) prepare() {

}

func New(paymentMethod string, address string, itens []Item, merchant Merchant) (*Order, error) {
	duuid := GenerateNewUUID()
	order := &Order{
		Id:            duuid,
		PaymentMethod: paymentMethod,
		Status:        Status{kind: Pending, reason: ""},
		Address:       address,
		Itens:         itens,
		Merchant:      merchant,
	}

	return order, nil
}
