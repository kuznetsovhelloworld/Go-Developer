package main

import (
	"errors"
	"unicode"
)

// CheckPasswordComplexity validates a password string against specific criteria.
// This function enforces stronger password rules to enhance security.
// Criteria enforced:
// 1. Minimum length of 8 characters
// 2. Contains at least one digit
// 3. Contains at least one special character (symbol or punctuation)
func CheckPasswordComplexity(password string) error {
	// Check length first to ensure password isn't too short.
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Track the presence of required character types.
	hasDigit := false
	hasSpecial := false

	// Iterate through each character in the password to identify digits and special characters.
	for _, c := range password {
		switch {
		case unicode.IsDigit(c):
			// If character is a digit, mark that we found a digit.
			hasDigit = true
		case unicode.IsSymbol(c) || unicode.IsPunct(c):
			// If character is a symbol or punctuation, mark that we found a special character.
			hasSpecial = true
		}
	}

	// After checking all characters, ensure that both digit and special character requirements are met.
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	// If all checks pass, the password meets the complexity requirements.
	return nil
}
