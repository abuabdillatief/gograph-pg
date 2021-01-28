package resolvers
//go:generate go run github.com/99designs/gqlgen
// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"github.com/abuabdillatief/gograph-tutorial/database"
	"github.com/abuabdillatief/gograph-tutorial/graph/generated"
	"github.com/abuabdillatief/gograph-tutorial/graph/model"
)

var meetups = []*model.Meetup{
	{
		ID:          "1",
		Name:        "A meetup",
		Description: "A Description",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "A second meetup",
		Description: "A Description",
		UserID:      "2",
	},
}

//Resolver ...
type Resolver struct {
	MeetupsRepo database.MeetupsRepo
	UsersRepo   database.UsersRepo
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
