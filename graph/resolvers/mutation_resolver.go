package resolvers

import (
	"context"
	"errors"

	"github.com/abuabdillatief/gograph-tutorial/graph"
	"github.com/abuabdillatief/gograph-tutorial/graph/model"
)

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetupInput) (*model.Meetup, error) {
	return r.Domain.CreateMeetup(ctx, input)
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetupInput) (*model.Meetup, error) {
	return r.Domain.UpdateMeetup(ctx, id, input)
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (*model.Response, error) {
	return r.Domain.DeleteMeetup(ctx, id)
}

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	isValid := graph.ValidationChecker(ctx, input)
	if !isValid {
		return nil, errors.New("input error")
	}
	return r.Domain.Register(ctx, input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	isValid := graph.ValidationChecker(ctx, input)
	if !isValid {
		return nil, errors.New("input error")
	}
	return r.Domain.Login(ctx, input)
}
