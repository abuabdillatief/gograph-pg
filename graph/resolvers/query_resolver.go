package resolvers

import (
	"context"

	"github.com/abuabdillatief/gograph-tutorial/graph/model"
)

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	return r.MeetupsRepo.GetMeetups()
}
