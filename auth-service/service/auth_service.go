package service

import (
    "auth-service/models"
    "auth-service/repository"
    "errors"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
)

type AuthService struct {
    repo      *repository.AuthRepository
    jwtSecret string
}

func NewAuthService(repo *repository.AuthRepository, jwtSecret string) *AuthService {
    return &AuthService{repo: repo, jwtSecret: jwtSecret}
}

func (s *AuthService) Register(username, password, email string) error {
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

func (s *AuthService) Login(username, password string) (string, error) {
    user, err := s.repo.GetUserByUsername(username)
    if err != nil || user == nil {
        return "", errors.New("invalid username or password")
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return "", errors.New("invalid username or password")
    }

    token, err := s.generateJWT(user)
    if err != nil {
        return "", err
    }

    return token, nil
}

func (s *AuthService) generateJWT(user *models.User) (string, error) {
    claims := jwt.MapClaims{
        "userID": user.ID,
        "exp":    time.Now().Add(time.Hour * 72).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.jwtSecret))
}
