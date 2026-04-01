package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"errors"
	"belajar-go/models"
	"time"
)

func GenerateJWT(user *models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT secret is not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"email":   user.Email,
		"name":    user.Name,
		"exp":	 jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		"iat":	 jwt.NewNumericDate(time.Now()),
	})

	return token.SignedString([]byte(secret))
}



func ValidateJWT(tokenString string) (int, error) {

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return 0, errors.New("JWT secret is not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				return 0, errors.New("token has expired")
			}
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdFloat, ok := claims["user_id"].(float64)
		if !ok {
			return 0, errors.New("invalid data in token")
		}
		return int(userIdFloat), nil
	} else {
		return 0, errors.New("invalid token")
	}
}