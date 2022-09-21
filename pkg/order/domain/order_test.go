package domain

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCancel(t *testing.T) {
	t.Run("only changes a order statuses to canceled when is different from preparing or delivered", func(t *testing.T) {
		itens := []Item{{quantity: 1, product: Product{id: "udfishuifhas", name: "café"}, comment: "sem açúcar"}}
		merchant := Merchant{id: "uioauiogugi", name: "Test", address: "Rua da tristeza"}
		expectedStatus := Status{kind: Canceled, reason: "comida fria"}

		order, _ := New("credit", "rua das maravilhas", itens, merchant)

		order.cancel("comida fria")

		assert.Equal(t, expectedStatus, order.Status)
	})

	t.Run("returns an error when try to cancel an order that was delivered ", func(t *testing.T) {
		itens := []Item{{quantity: 1, product: Product{id: "udfishuifhas", name: "café"}, comment: "sem açúcar"}}
		merchant := Merchant{id: "uioauiogugi", name: "Test", address: "Rua da tristeza"}
		expectedError := errors.New("cannot cancel the order")

		order, _ := New("credit", "rua das maravilhas", itens, merchant)
		order.delivered()

		err := order.cancel("estou maluco")

		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
	})
}

func generateUUID(uuid string) func() string {
	GenerateNewUUID = func() string { return uuid }
	return GenerateNewUUID
}
