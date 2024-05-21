package main

import (
	"api-gateway/config"
	"api-gateway/handlers"
	"api-gateway/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	r := mux.NewRouter()

	// Middleware
	r.Use(middlewares.AuthMiddleware)

	// Cart routes
	r.HandleFunc("/cart", handlers.AddToCart).Methods("POST")
	r.HandleFunc("/cart/{userId}", handlers.GetCart).Methods("GET")
	r.HandleFunc("/cart/{userId}/item/{itemId}", handlers.RemoveFromCart).Methods("DELETE")

	// Payment routes
	r.HandleFunc("/payments", handlers.CreatePayment).Methods("POST")
	r.HandleFunc("/payments/{id}", handlers.GetPayment).Methods("GET")
	r.HandleFunc("/payments/{id}", handlers.UpdatePayment).Methods("PUT")
	r.HandleFunc("/payments/{id}", handlers.DeletePayment).Methods("DELETE")

	// Order routes
	r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", handlers.GetOrder).Methods("GET")
	r.HandleFunc("/orders/{id}", handlers.UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", handlers.DeleteOrder).Methods("DELETE")

	// User routes
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	log.Printf("API Gateway running on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
