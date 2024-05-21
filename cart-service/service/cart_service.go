package service

import (
    "cart-service/models"
    "cart-service/repository"
)

type CartService struct {
    repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
    return &CartService{repo: repo}
}

func (s *CartService) AddToCart(userID, productID, quantity int) error {
    item := &models.CartItem{
        UserID:    userID,
        ProductID: productID,
        Quantity:  quantity,
    }

    return s.repo.AddToCart(item)
}

func (s *CartService) GetCart(userID int) ([]*models.CartItem, error) {
    return s.repo.GetCart(userID)
}

func (s *CartService) RemoveFromCart(userID, itemID int) error {
    return s.repo.RemoveFromCart(userID, itemID)
}
