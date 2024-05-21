package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "auth-service/config"
    "auth-service/handlers"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Could not load config: %v", err)
    }

    r := mux.NewRouter()

    // Auth routes
    r.HandleFunc("/login", handlers.Login).Methods("POST")
    r.HandleFunc("/register", handlers.Register).Methods("POST")

    log.Printf("Auth Service running on port %s", cfg.Port)
    if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}
