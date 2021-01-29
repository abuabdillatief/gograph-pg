package validator

//Validation ...
type Validation interface {
	Validate() (bool, map[string]string)
}

//Validator ...
type Validator struct {
	Errors map[string]string
}

//New ...
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

//IsValid ...
func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}
