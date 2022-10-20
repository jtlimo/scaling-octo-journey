package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"order/pkg/order/application"
	"order/pkg/order/domain"
	"order/pkg/order/repository"
	"order/pkg/order/routes/adapters"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var s = Server{
	Router:      mux.NewRouter(),
	Application: application.New((repository.New())),
}

func TestCreateOrder(t *testing.T) {
	newServer(t)
	var old = domain.GenerateNewUUID
	defer func() { domain.GenerateNewUUID = old }()

	t.Run("create an order", func(t *testing.T) {
		expectedDeck := adapters.OrderResponse{
			Id:            "0cf727b0-aaae-499d-9f12-e9c6dcee1d45",
			PaymentMethod: "Jess",
			Status:        "pending",
			Address:       "Rua uruguai, 206",
			Itens: []adapters.Item{{
				Quantity: 1,
				Product:  adapters.Product{Id: "1"},
				Comment:  "que delicia",
			}, {
				Quantity: 10,
				Product:  adapters.Product{Id: "2"},
				Comment:  "que podre",
			}},
			Merchant: adapters.Merchant{Id: "1000"},
		}
		generateUUID(t, "0cf727b0-aaae-499d-9f12-e9c6dcee1d45")()

		body := adapters.OrderRequestBody{
			Itens: []adapters.Item{{
				Quantity: 1,
				Product:  adapters.Product{Id: "1"},
				Comment:  "que delicia",
			}, {
				Quantity: 10,
				Product:  adapters.Product{Id: "2"},
				Comment:  "que podre",
			}},
			PaymentMethod: "Jess",
			Address:       "Rua uruguai, 206",
		}

		request := createAnOrder("1000", body)
		res := executeRequest(request)

		payload, _ := ioutil.ReadAll(res.Body)

		var jsonData adapters.OrderResponse
		json.Unmarshal(payload, &jsonData)

		if assert.NotNil(t, jsonData) {
			assert.Equal(t, expectedDeck, jsonData)
		}
		assertStatus(t, res.Code, http.StatusOK)
	})
}

func createAnOrder(companyId string, payload adapters.OrderRequestBody) *http.Request {
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/company/%s/order", companyId), bytes.NewBuffer(body))

	return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func newServer(t *testing.T) {
	t.Helper()
	localServer := httptest.NewServer(s.Router)
	s.Register()
	defer localServer.Close()
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func generateUUID(t *testing.T, uuid string) func() string {
	t.Helper()
	domain.GenerateNewUUID = func() string { return uuid }
	return domain.GenerateNewUUID
}
