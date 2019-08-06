package utils

import (
	"regexp"
)

const (
	emailRegex = `(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`
)

// ValidateEmail checks wheter or not an email is valid
func ValidateEmail(email string) bool {
	match, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return false
	}

	return match
}

// ValidatePassword checks wheter or not a password is valid
func ValidatePassword(password string) bool {
	return len(password) >= 8
}
