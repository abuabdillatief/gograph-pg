package validator

import "fmt"

//EqualToField ...
func (v *Validator) EqualToField(field, toEqualField string, value, toEqualValue interface{}) bool {
	if _, notOK := v.Errors[field]; notOK {
		return false
	}
	if value != toEqualValue {
		v.Errors[field] = fmt.Sprintf("%s must equal %s", field, toEqualField)
	}
	return true
}
