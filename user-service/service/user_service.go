package service

import (
    "user-service/models"
    "user-service/repository"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(username, password, email string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := &models.User{
        Username: username,
        Password: string(hashedPassword),
        Email:    email,
    }

    return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
    return s.repo.GetUserByID(id)
}

func (s *UserService) UpdateUser(id int, username, password, email string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := &models.User{
        ID:       id,
        Username: username,
        Password: string(hashedPassword),
        Email:    email,
    }

    return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
    return s.repo.DeleteUser(id)
}
