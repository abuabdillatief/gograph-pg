package resolvers

import (
	"context"

	"github.com/abuabdillatief/gograph-tutorial/graph/model"
)

//this means, this resolver is for meetup queries which specifies on User object.
//meaning any queries that mainly fetch meetups but also has User as child object
//thats why its called meetupResolver BUT IN THAT, still queries User.
func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	//.Load(obj.UserID) here is defined in generated dataloaden code
	//it is used to load a user by key
	return model.GetUserLoader(ctx).Load(obj.UserID)
}
