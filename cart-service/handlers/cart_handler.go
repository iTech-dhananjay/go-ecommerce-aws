package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "cart-service/config"
    "cart-service/models"
    "cart-service/repository"
    "cart-service/service"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

var (
    cfg      *config.Config
    cartRepo *repository.CartRepository
    cartSvc  *service.CartService
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

    cartRepo = repository.NewCartRepository(db)
    cartSvc = service.NewCartService(cartRepo)
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
    var req struct {
        UserID    int `json:"user_id"`
        ProductID int `json:"product_id"`
        Quantity  int `json:"quantity"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := cartSvc.AddToCart(req.UserID, req.ProductID, req.Quantity); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func GetCart(w http.ResponseWriter, r *http.Request) {
    userID, err := strconv.Atoi(mux.Vars(r)["userId"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    items, err := cartSvc.GetCart(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if items == nil {
        http.Error(w, "No items found in cart", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(items)
}

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
    userID, err := strconv.Atoi(mux.Vars(r)["userId"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    itemID, err := strconv.Atoi(mux.Vars(r)["itemId"])
    if err != nil {
        http.Error(w, "Invalid item ID", http.StatusBadRequest)
        return
    }

    if err := cartSvc.RemoveFromCart(userID, itemID); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
