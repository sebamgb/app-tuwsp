package handlers

import (
	"unicode"
)

func validatePassword(password string) (valid bool) {
	if len(password) < 8 {
		return
	}
	for _, char := range password {
		if unicode.IsNumber(char) {
			valid = true
			break
		}
	}
	return
}
