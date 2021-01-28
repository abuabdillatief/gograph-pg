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

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*model.Meetup, error) {
	meetup, err := r.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return nil, fmt.Errorf("meetup no exist")
	}
	updated := false
	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, fmt.Errorf("name is not long enough")
		}
		meetup.Name = *input.Name
		updated = true

	}
	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, fmt.Errorf("description is not long enough")
		}
		meetup.Description = *input.Description
		updated = true
	}
	if !updated {
		return nil, fmt.Errorf("no update is applied")
	}
	meetup, err = r.MeetupsRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updating meetup object: %v", err)
	}
	return meetup, nil
}
