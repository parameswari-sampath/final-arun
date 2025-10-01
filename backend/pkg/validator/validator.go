package validator

import (
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func IsValidPassword(password string) bool {
	// Minimum 8 characters
	return len(password) >= 8
}

func IsValidName(name string) bool {
	return len(name) >= 2
}
