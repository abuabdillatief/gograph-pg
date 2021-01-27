package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/abuabdillatief/gograph-tutorial/graph/generated"
	"github.com/abuabdillatief/gograph-tutorial/graph/model"
)

type Resolver struct{}

func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreaetMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	panic("not implemented")
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	panic("not implemented")
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
