package middleware

import (
	"Flowershop-GoBackend/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

func HashwordCompare(hashword string, user models.User) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashword), []byte(user.Password))
	return err
}
