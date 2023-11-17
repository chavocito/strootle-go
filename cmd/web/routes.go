package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	//Public Routes
	mux.Group(func(r chi.Router) {
		mux.Get("/virtual-terminal", app.VirtualTerminal)
		mux.Post("/payment-information/{userId}", app.PaymentConfirmation)
	})

	return mux
}
