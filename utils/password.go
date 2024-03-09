package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword mengembalikan hash kata sandi yg di encrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword memeriksa apakah kata sandi yang diberikan benar atau tidak
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
