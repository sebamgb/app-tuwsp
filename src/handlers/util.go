package handlers

import (
	"regexp"
	"unicode"
)

// validateEmail do validation of an email
func validateEmail(email string) (valid bool) {
	valid = regexp.
		MustCompile(
			"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
		).MatchString(email)
	return
}

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
