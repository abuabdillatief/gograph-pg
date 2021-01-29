package domain

import (
	"context"
	"fmt"

	"github.com/abuabdillatief/gograph-tutorial/graph/model"
	"github.com/abuabdillatief/gograph-tutorial/middlewares"
)

//CreateMeetup ...
func (d *Domain) CreateMeetup(ctx context.Context, input model.NewMeetupInput) (*model.Meetup, error) {
	currentUser, err := middlewares.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if len(input.Name) < 3 {
		return nil, fmt.Errorf("name is not long enough")
	}
	if len(input.Description) < 3 {
		return nil, fmt.Errorf("description is not long enough")
	}
	meetup := &model.Meetup{Name: input.Name,
		Description: input.Description,
		UserID:      currentUser.ID}
	return d.MeetupsRepo.CreateMeetup(meetup)
}

//UpdateMeetup ...
func (d *Domain) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetupInput) (*model.Meetup, error) {
	meetup, err := d.MeetupsRepo.GetByID(id)
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
	meetup, err = d.MeetupsRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updating meetup object: %v", err)
	}
	return meetup, nil
}

//DeleteMeetup ...
func (d *Domain) DeleteMeetup(ctx context.Context, id string) (*model.Response, error) {
	var res model.Response
	meetup, err := d.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		res.Message = fmt.Sprintf("meetup with id %v is not found", id)
		return &res, nil
	}
	err = d.MeetupsRepo.Delete(id)
	if err != nil {
		return nil, fmt.Errorf("cant delete object with id %v", id)
	}
	res.Message = fmt.Sprintf("meetup with id %v has been deleted", id)
	return &res, nil

}
