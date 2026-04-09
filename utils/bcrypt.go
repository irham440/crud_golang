package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

func GeneratePassword(password string) (string, error) {
	cost := os.Getenv("COST")
	costValue, err := strconv.Atoi(cost)
	if err != nil {
		return "", err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), costValue)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedpassword []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedpassword, password)
	if err != nil {
		return errors.New("invalid email or password")
	}
	return nil
}
