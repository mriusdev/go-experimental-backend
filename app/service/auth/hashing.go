package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) string {
	bytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error while hashing", err.Error())
	}

	return string(bytes)
}

func VerifyPassword(password []byte, hashedPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		fmt.Println("Error while verifying password", err.Error())
		return false
	}

	return true
}
