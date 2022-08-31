package domain

// FIXME: change int to uuid for id's
type Order struct {
	id            int
	paymentMethod string
	status        Status
	address       string
	itens         []Item
	merchant      Merchant
}

type Status struct {
	kind   string
	reason string
}

type Item struct {
	quantity int
	product  Product
	comment  string
}

type Product struct {
	id   int
	name string
}

type Merchant struct {
	id      int
	name    string
	address string
}

// TODO: verify status before cancelling the order
func (o *Order) cancelOrder(reason string) {
	o.status.kind = "canceled"
	o.status.reason = reason
}

// TODO: create a method to prepare a order
