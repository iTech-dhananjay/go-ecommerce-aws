package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "payment-service/config"
    "payment-service/models"
    "payment-service/repository"
    "payment-service/service"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

var (
    cfg       *config.Config
    paymentRepo *repository.PaymentRepository
    paymentSvc  *service.PaymentService
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

    paymentRepo = repository.NewPaymentRepository(db)
    paymentSvc = service.NewPaymentService(paymentRepo)
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
    var req struct {
        OrderID int    `json:"order_id"`
        Amount  int    `json:"amount"`
        Status  string `json:"status"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := paymentSvc.CreatePayment(req.OrderID, req.Amount, req.Status); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid payment ID", http.StatusBadRequest)
        return
    }

    payment, err := paymentSvc.GetPaymentByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if payment == nil {
        http.Error(w, "Payment not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(payment)
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid payment ID", http.StatusBadRequest)
        return
    }

    var req struct {
        OrderID int    `json:"order_id"`
        Amount  int    `json:"amount"`
        Status  string `json:"status"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := paymentSvc.UpdatePayment(id, req.OrderID, req.Amount, req.Status); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func DeletePayment(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid payment ID", http.StatusBadRequest)
        return
    }

    if err := paymentSvc.DeletePayment(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
