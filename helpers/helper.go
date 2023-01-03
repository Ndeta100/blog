package helpers

import (
	"errors"
	"regexp"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Validate email base on regex
func ValidateEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return errors.New("invalid string format")
	}
	return nil
}
