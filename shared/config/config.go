package config

import (
    "github.com/joho/godotenv"
    "os"
)

type Config struct {
    AuthServiceURL string
    CartServiceURL string
    OrderServiceURL string
    PaymentServiceURL string
    UserServiceURL string
    JWTSecret       string
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    cfg := &Config{
        AuthServiceURL:  os.Getenv("AUTH_SERVICE_URL"),
        CartServiceURL:  os.Getenv("CART_SERVICE_URL"),
        OrderServiceURL: os.Getenv("ORDER_SERVICE_URL"),
        PaymentServiceURL: os.Getenv("PAYMENT_SERVICE_URL"),
        UserServiceURL:  os.Getenv("USER_SERVICE_URL"),
        JWTSecret:       os.Getenv("JWT_SECRET"),
    }

    return cfg, nil
}
