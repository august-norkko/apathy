package response

import (
	"regexp"
)

func ValidateUser(email, password string) bool { // move from response package
	ok, err := regexp.MatchString(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`, email)
	if err != nil {
		return false
	}

	if !ok {
		return false
	}

	if len(password) < 5 {
		return false
	}

	return true
}