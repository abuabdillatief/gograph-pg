package validation

//Validator ...
type Validator interface {
	Validate() (bool, map[string]string)
}

//Validation ...
type Validation struct {
	Errors map[string]string
}

//NewValidation ...
func NewValidation() *Validation {
	return &Validation{Errors: make(map[string]string)}
}

//IsValid ...
func (v *Validation) IsValid() bool {
	return len(v.Errors) == 0
}
