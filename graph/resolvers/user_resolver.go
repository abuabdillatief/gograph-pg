package resolvers

import (
	"context"

	"github.com/abuabdillatief/gograph-tutorial/graph/model"
)

func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	var m []*model.Meetup
	for _, meetup := range meetups {
		if meetup.UserID == obj.ID {
			m = append(m, meetup)
		}
	}
	return meetups, nil
}
