package resolvers

import (
	"context"

	"github.com/abuabdillatief/gograph-tutorial/graph/model"
)

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	return r.MeetupsRepo.GetMeetups()
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.UsersRepo.GerUserByID(id)
}
