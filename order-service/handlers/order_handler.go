package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "order-service/config"
    "order-service/models"
    "order-service/repository"
    "order-service/service"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

var (
    cfg       *config.Config
    orderRepo *repository.OrderRepository
    orderSvc  *service.OrderService
)

func init() {
    var err error
    cfg, err = config.LoadConfig()
    if err != nil {
        panic(err)
    }

    db, err := sqlx.Connect("postgres", "host="+cfg.DBHost+" port="+cfg.DBPort+" user="+cfg.DBUser+" password="+cfg.DBPassword+" dbname="+cfg.DBName+" sslmode=disable")
    if err != nil {
        panic(err)
    }

    orderRepo = repository.NewOrderRepository(db)
    orderSvc = service.NewOrderService(orderRepo)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
    var req struct {
        UserID    int    `json:"user_id"`
        ProductID int    `json:"product_id"`
        Quantity  int    `json:"quantity"`
        Status    string `json:"status"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := orderSvc.CreateOrder(req.UserID, req.ProductID, req.Quantity, req.Status); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid order ID", http.StatusBadRequest)
        return
    }

    order, err := orderSvc.GetOrderByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if order == nil {
        http.Error(w, "Order not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(order)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid order ID", http.StatusBadRequest)
        return
    }

    var req struct {
        UserID    int    `json:"user_id"`
        ProductID int    `json:"product_id"`
        Quantity  int    `json:"quantity"`
        Status    string `json:"status"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := orderSvc.UpdateOrder(id, req.UserID, req.ProductID, req.Quantity, req.Status); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid order ID", http.StatusBadRequest)
        return
    }

    if err := orderSvc.DeleteOrder(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
