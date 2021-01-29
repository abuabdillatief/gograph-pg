package validator

import "regexp"

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

//IsEmail ...
func (v *Validator) IsEmail(field, email string) bool {
	if _, notOK := v.Errors[field]; notOK {
		return false
	}

	if !emailRegex.MatchString(email) {
		v.Errors[field] = "Email format is not valid"
		return false
	}
	return true
}
