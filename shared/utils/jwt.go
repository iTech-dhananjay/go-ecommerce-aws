package utils

import (
    "errors"
    "github.com/dgrijalva/jwt-go"
    "time"
)

func GenerateJWT(userID int, secret string) (string, error) {
    claims := jwt.MapClaims{
        "userID": userID,
        "exp":    time.Now().Add(time.Hour * 72).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

func ValidateToken(tokenStr string, secret string) (bool, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(secret), nil
    })

    if err != nil {
        return false, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if exp, ok := claims["exp"].(float64); ok {
            if int64(exp) < time.Now().Unix() {
                return false, errors.New("token expired")
            }
            return true, nil
        }
    }

    return false, errors.New("invalid token")
}
