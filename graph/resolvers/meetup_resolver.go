package resolvers

import (
	"context"

	"github.com/abuabdillatief/gograph-tutorial/graph/model"
)

func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	return model.GetUserLoader(ctx).Load(obj.UserID)
}
