package model

import (
	"github.com/abuabdillatief/gograph-tutorial/validation"
)

//Validate ...
func (r RegisterInput) Validate() (bool, map[string]string) {
	v := validation.NewValidation()
	v.IsRequired("email", r.Email)
	v.IsEmail("email", r.Email)

	v.IsRequired("password", r.Password)
	v.MinLength("password", r.Password, 6)

	v.IsRequired("confirmPassword", r.ConfirmPassword)
	v.EqualToField("confirmPassword", "password", r.ConfirmPassword, r.Password)

	v.IsRequired("username", r.Username)
	v.MinLength("username", r.Username, 6)

	v.IsRequired("firstName", r.FirstName)
	v.MinLength("firstName", r.FirstName, 6)

	v.IsRequired("lastName", r.LastName)
	v.MinLength("lastName", r.LastName, 6)
	return v.IsValid(), v.Errors
}

//Validate ...
func (l LoginInput) Validate() (bool, map[string]string) {
	v := validation.NewValidation()
	v.IsRequired("email", l.Email)
	v.IsEmail("email", l.Email)

	v.IsRequired("password", l.Password)
	return v.IsValid(), v.Errors
}
