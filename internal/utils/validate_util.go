package utils

import (
	"authenncrm/internal/constants"
	"authenncrm/internal/entities"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(user *entities.UserRequest) string {

	if user.Password != user.ConfirmPassword {
		return constants.PasswordMismatchError
	}

	if len(user.Password) < 8 {
		return constants.PasswordLengthError
	}

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	for _, char := range user.Password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasDigit = true
		case (char >= 33 && char <= 47) || (char >= 58 && char <= 64) ||
			(char >= 91 && char <= 96) || (char >= 123 && char <= 126):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
		return constants.PasswordComplexityError
	}

	return constants.PassedValidation
}

func CheckPasswordEncryption(password string, encryptedPassword string) bool {

	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
