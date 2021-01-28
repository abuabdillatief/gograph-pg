package resolvers

import (
	"context"

	"github.com/abuabdillatief/gograph-tutorial/graph/model"
)

func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	return r.MeetupsRepo.GetSelectedMeetups(obj.ID)
}
