package handler

import (
	"encoding/json"
	"net/http"

	"github.com/telmocbarros/cozy-wallet/internal/handler/dto"
)

type PaymentCardHandler struct{}

func (pc PaymentCardHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Payment Cards"))
}

func (pc PaymentCardHandler) GetSingle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Single Payment Card"))
}

func (pc PaymentCardHandler) Create(w http.ResponseWriter, r *http.Request) {
	var paymentCard dto.CreateCardRequest
	err := json.NewDecoder(r.Body).Decode(&paymentCard)
	if err != nil {
		w.Write([]byte("Create Payment Card"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Payment card created"))
}

func (pc PaymentCardHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Payment Card"))
}

func (pc PaymentCardHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Payment Card"))
}
