package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "payment-service/config"
    "payment-service/handlers"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Could not load config: %v", err)
    }

    r := mux.NewRouter()

    // Payment routes
    r.HandleFunc("/payments", handlers.CreatePayment).Methods("POST")
    r.HandleFunc("/payments/{id}", handlers.GetPayment).Methods("GET")
    r.HandleFunc("/payments/{id}", handlers.UpdatePayment).Methods("PUT")
    r.HandleFunc("/payments/{id}", handlers.DeletePayment).Methods("DELETE")

    log.Printf("Payment Service running on port %s", cfg.Port)
    if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}
