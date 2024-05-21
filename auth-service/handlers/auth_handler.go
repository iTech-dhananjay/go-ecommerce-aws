package handlers

import (
    "auth-service/config"
    "auth-service/repository"
    "auth-service/service"
    "encoding/json"
    "net/http"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

var (
    cfg       *config.Config
    authRepo  *repository.AuthRepository
    authSvc   *service.AuthService
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

    authRepo = repository.NewAuthRepository(db)
    authSvc = service.NewAuthService(authRepo, cfg.JWTSecret)
}

func Register(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
        Email    string `json:"email"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := authSvc.Register(req.Username, req.Password, req.Email); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    token, err := authSvc.Login(req.Username, req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.Header().Set("Authorization", "Bearer "+token)
    w.WriteHeader(http.StatusOK)
}
