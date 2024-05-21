package service

import (
    "payment-service/models"
    "payment-service/repository"
)

type PaymentService struct {
    repo *repository.PaymentRepository
}

func NewPaymentService(repo *repository.PaymentRepository) *PaymentService {
    return &PaymentService{repo: repo}
}

func (s *PaymentService) CreatePayment(orderID, amount int, status string) error {
    payment := &models.Payment{
        OrderID: orderID,
        Amount:  amount,
        Status:  status,
    }

    return s.repo.CreatePayment(payment)
}

func (s *PaymentService) GetPaymentByID(id int) (*models.Payment, error) {
    return s.repo.GetPaymentByID(id)
}

func (s *PaymentService) UpdatePayment(id, orderID, amount int, status string) error {
    payment := &models.Payment{
        ID:       id,
        OrderID:  orderID,
        Amount:   amount,
        Status:   status,
    }

    return s.repo.UpdatePayment(payment)
}

func (s *PaymentService) DeletePayment(id int) error {
    return s.repo.DeletePayment(id)
}
