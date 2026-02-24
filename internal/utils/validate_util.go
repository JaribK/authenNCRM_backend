package utils

import (
	"authenncrm/internal/constants"
	"authenncrm/internal/entities"
)

func ValidatePassword(user *entities.UserRequest) string {

	//check confirm password
	if user.Password != user.ConfirmPassword {
		return constants.PasswordMismatchError
	}

	//check password length
	if len(user.Password) < 8 {
		return constants.PasswordLengthError
	}

	//check password complexity
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
