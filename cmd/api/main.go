package main

import (
	"net/http"

	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/telmocbarros/cozy-wallet/internal/handler"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/payment-cards", PaymentCardRouter())

	fmt.Println("Server listening on port 3000")
	http.ListenAndServe(":3000", r)
}

func PaymentCardRouter() chi.Router {
	r := chi.NewRouter()
	paymentCardHandler := handler.PaymentCardHandler{}
	r.Post("/", paymentCardHandler.Create)
	r.Get("/", paymentCardHandler.GetAll)
	r.Get("/{id}", paymentCardHandler.GetSingle)
	r.Put("/{id}", paymentCardHandler.Update)
	r.Delete("/{id}", paymentCardHandler.Delete)
	return r
}
