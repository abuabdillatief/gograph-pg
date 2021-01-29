package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/abuabdillatief/gograph-tutorial/validator"
	"github.com/vektah/gqlparser/gqlerror"
)

func validation(ctx context.Context, v validator.Validation) bool {
	isValid, err := v.Validate()
	if !isValid {
		for key, err := range err {
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
 