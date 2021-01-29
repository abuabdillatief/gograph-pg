package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	customValidation "github.com/abuabdillatief/gograph-tutorial/validation"
	"github.com/vektah/gqlparser/gqlerror"
)

//ValidationChecker ...
func ValidationChecker(ctx context.Context, validator customValidation.Validator) bool {
	isValid, errors := validator.Validate()
	if !isValid {
		for key, err := range errors {
			graphql.AddError(ctx, &gqlerror.Error{
				Message: err,
				Extensions: map[string]interface{}{
					"field": key,
				},
			})
		}
	}
	return isValid
}
