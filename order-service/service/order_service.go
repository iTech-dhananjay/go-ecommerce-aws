package service

import (
    "order-service/models"
    "order-service/repository"
)

type OrderService struct {
    repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
    return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(userID, productID, quantity int, status string) error {
    order := &models.Order{
        UserID:    userID,
        ProductID: productID,
        Quantity:  quantity,
        Status:    status,
    }

    return s.repo.CreateOrder(order)
}

func (s *OrderService) GetOrderByID(id int) (*models.Order, error) {
    return s.repo.GetOrderByID(id)
}

func (s *OrderService) UpdateOrder(id, userID, productID, quantity int, status string) error {
    order := &models.Order{
        ID:        id,
        UserID:    userID,
        ProductID: productID,
        Quantity:  quantity,
        Status:    status,
    }

    return s.repo.UpdateOrder(order)
}

func (s *OrderService) DeleteOrder(id int) error {
    return s.repo.DeleteOrder(id)
}
