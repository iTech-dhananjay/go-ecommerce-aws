package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	CartServiceURL    string
	PaymentServiceURL string
	OrderServiceURL   string
	UserServiceURL    string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Port:              os.Getenv("API_GATEWAY_PORT"),
		CartServiceURL:    os.Getenv("CART_SERVICE_URL"),
		PaymentServiceURL: os.Getenv("PAYMENT_SERVICE_URL"),
		OrderServiceURL:   os.Getenv("ORDER_SERVICE_URL"),
		UserServiceURL:    os.Getenv("USER_SERVICE_URL"),
	}

	return cfg, nil
}
