package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "order-service/config"
    "order-service/handlers"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Could not load config: %v", err)
    }

    r := mux.NewRouter()

    // Order routes
    r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
    r.HandleFunc("/orders/{id}", handlers.GetOrder).Methods("GET")
    r.HandleFunc("/orders/{id}", handlers.UpdateOrder).Methods("PUT")
    r.HandleFunc("/orders/{id}", handlers.DeleteOrder).Methods("DELETE")

    log.Printf("Order Service running on port %s", cfg.Port)
    if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}
