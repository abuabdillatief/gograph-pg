package validator

import (
	"fmt"
	"reflect"
)

//IsEmpty ...
func (v *Validator) IsEmpty(value interface{}) bool {
	t := reflect.ValueOf(value)
	switch t.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return t.Len() == 0
	}
	return false
}

//Required ...
func (v *Validator) Required(field string, value interface{}) bool {
	if _, notOK := v.Errors[field]; notOK {
		return false
	}
	if v.IsEmpty(value) {
		v.Errors[field] = fmt.Sprintf("%s is required", field)
		return false
	}
	return true
}
