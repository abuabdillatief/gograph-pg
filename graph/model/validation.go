package model

import "github.com/abuabdillatief/gograph-tutorial/validator"

//Validate ...
func (r RegisterInput) Validate() (bool, map[string]string) {
	v := validator.New()
	v.Required("email", r.Email)
	v.IsEmail("email", r.Email)
	v.Required("password", r.Password)
	v.MinLength("password", r.Password, 6)
	v.Required("confirmPassword", r.ConfirmPassword)
	v.EqualToField("confirmPassword", "password", r.ConfirmPassword, r.Password)
	v.Required("username", r.Username)
	v.MinLength("username", r.Username, 6)
	return v.IsValid(), v.Errors
}

//Validate ...
func (l LoginInput) Validate() (bool, map[string]string) {
	v := validator.New()
	v.Required("email", l.Email)
	v.IsEmail("email", l.Email)
	v.Required("password", l.Password)
	return v.IsValid(), v.Errors
}
