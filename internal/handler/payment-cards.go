package handler

import (
	"net/http"
)

type PaymentCardHandler struct{}

func (pc PaymentCardHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Payment Cards"))
}

func (pc PaymentCardHandler) GetSingle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Single Payment Card"))
}

func (pc PaymentCardHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Payment Card"))
}

func (pc PaymentCardHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Payment Card"))
}

func (pc PaymentCardHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Payment Card"))
}
