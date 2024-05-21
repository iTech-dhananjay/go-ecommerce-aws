package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "cart-service/config"
    "cart-service/handlers"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Could not load config: %v", err)
    }

    r := mux.NewRouter()

    // Cart routes
    r.HandleFunc("/cart", handlers.AddToCart).Methods("POST")
    r.HandleFunc("/cart/{userId}", handlers.GetCart).Methods("GET")
    r.HandleFunc("/cart/{userId}/item/{itemId}", handlers.RemoveFromCart).Methods("DELETE")

    log.Printf("Cart Service running on port %s", cfg.Port)
    if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}
