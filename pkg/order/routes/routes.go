package routes

import (
	"encoding/json"
	"net/http"
	"order/pkg/order/domain"
	"order/pkg/order/routes/adapters"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (s *Server) Register() {
	s.Router.HandleFunc("/company/{id}/order", createOrder).Methods("POST")
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var orderRequest adapters.OrderRequestBody
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		http.Error(w, "Error decoding response object", http.StatusBadRequest)
		return
	}

	_, err := json.Marshal(&orderRequest)
	if err != nil {
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}

	o := adapters.AdaptToDomain(vars["id"], orderRequest)
	orderDomain, err := domain.New(o.PaymentMethod, o.Address, o.Itens, o.Merchant)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	orderPayload, err := json.Marshal(adapters.Adapt(orderDomain))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write(orderPayload)
}
