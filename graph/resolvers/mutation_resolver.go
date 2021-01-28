package resolvers

import (
	"context"
	"fmt"

	"github.com/abuabdillatief/gograph-tutorial/graph/model"
)

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, fmt.Errorf("name is not long enough")
	}
	if len(input.Description) < 3 {
		return nil, fmt.Errorf("description is not long enough")
	}

	meetup := &model.Meetup{Name: input.Name,
		Description: input.Description,
		UserID:      "1"}
	return r.MeetupsRepo.CreateMeetup(meetup)
}
