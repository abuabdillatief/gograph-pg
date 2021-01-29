package validation

import "fmt"

//MinLength ...
func (v *Validation) MinLength(field, value string, minChar int) bool {
	if _, notOK := v.Errors[field]; notOK {
		return false
	}
	if len(value) < minChar {
		v.Errors[field] = fmt.Sprintf("%s must be at least %d characters", field, minChar)
		return false
	}
	return true
}
