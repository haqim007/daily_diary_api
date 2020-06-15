package helpers

import (
	"regexp"
)

// validation
type validation struct {
	string
}

//PasswordValidation a
func PasswordValidation(password string) bool {
	var maxLength = regexp.MustCompile("[\\s\\S]{8,}$")
	var upper = regexp.MustCompile("[A-Z]")
	var lower = regexp.MustCompile("[a-z]")
	var number = regexp.MustCompile("[0-9]")
	var special = regexp.MustCompile("[ !\"#$%&'()*+,\\-./:;<=>?@[\\]^_`{|}~]")

	// var re = regexp.MustCompile("^(?=.*[A-Za-z])(?=.*\\d)(?=.*[@$!%*#?&])[A-Za-z\\d@$!%*#?&]{8,16}$")
	if maxLength.MatchString(password) && upper.MatchString(password) && lower.MatchString(password) && number.MatchString(password) && special.MatchString(password) {
		return true
	}

	return false
}

//EmailValidation a
func EmailValidation(email string) bool {
	var re = regexp.MustCompile("^[A-Z0-9._%+-]+@[A-Z0-9.-]+\\.[A-Z]{2,}$")

	return re.MatchString(email)
}
