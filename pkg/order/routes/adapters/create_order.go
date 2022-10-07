package adapters

type OrderRequestBody struct {
	Itens         []Item `json:"itens"`
	MerchantId    string `json:"merchantId"`
	PaymentMethod string `json:"paymentMethod"`
	Address       string `json:"address"`
}

type Item struct {
	Quantity int     `json:"quantity"`
	Product  Product `json:"product"`
	Comment  string  `json:"comment"`
}

type Product struct {
	Id string `json:"id"`
}

type Kind string
type OrderResponse struct {
	Id            string   `json:"id"`
	PaymentMethod string   `json:"paymentMethod"`
	Status        Kind     `json:"status"`
	Address       string   `json:"address"`
	Itens         []Item   `json:"itens"`
	Merchant      Merchant `json:"merchant"`
}

type Merchant struct {
	Id string `json:"id"`
}
