package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "user-service/config"
    "user-service/handlers"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Could not load config: %v", err)
    }

    r := mux.NewRouter()

    // User routes
    r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

    log.Printf("User Service running on port %s", cfg.Port)
    if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}
