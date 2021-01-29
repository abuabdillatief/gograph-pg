package resolvers

import (
	"context"

	"github.com/abuabdillatief/gograph-tutorial/graph/model"
)

func (r *queryResolver) Meetups(ctx context.Context, filter *model.MeetupFilterInput, limit *int, offset *int) ([]*model.Meetup, error) {
	return r.MeetupsRepo.GetMeetups(filter, limit, offset)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.UsersRepo.GetUserByID(id)
}
