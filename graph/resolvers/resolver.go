package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"fmt"

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

func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	return r.UsersRepo.GerUserByID(obj.UserID)
}

func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	var m []*model.Meetup
	for _, meetup := range meetups {
		if meetup.UserID == obj.ID {
			m = append(m, meetup)
		}
	}
	return meetups, nil
}
func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, fmt.Errorf("Name is not long enough!")
	}
	if len(input.Description) < 3 {
		return nil, fmt.Errorf("Description is not long enough!")
	}

	meetup := &model.Meetup{Name: input.Name,
		Description: input.Description,
		UserID:      "1"}
	return r.MeetupsRepo.CreateMeetup(meetup)
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	return r.MeetupsRepo.GetMeetups()
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
