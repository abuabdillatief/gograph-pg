package resolvers

import (
	"context"
	"errors"
	"fmt"

	"github.com/abuabdillatief/gograph-tutorial/graph/model"
	"github.com/abuabdillatief/gograph-tutorial/middlewares"
)

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetupInput) (*model.Meetup, error) {
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
	return r.MeetupsRepo.CreateMeetup(meetup)
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetupInput) (*model.Meetup, error) {
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

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (*model.Response, error) {
	var res model.Response
	meetup, err := r.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		res.Message = fmt.Sprintf("meetup with id %v is not found", id)
		return &res, nil
	}
	err = r.MeetupsRepo.Delete(id)
	if err != nil {
		return nil, fmt.Errorf("cant delete object with id %v", id)
	}
	res.Message = fmt.Sprintf("meetup with id %v has been deleted", id)
	return &res, nil

}

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	_, err := r.UsersRepo.GetUserByEmail(input.Email)
	//we are expecting errors, as it shows that user doesn't exist
	if err == nil {
		return nil, errors.New("email is already in used")
	}
	_, err = r.UsersRepo.GetUserByUsername(input.Username)
	if err == nil {
		return nil, errors.New("username is already in used")
	}
	user := &model.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPass(input.Password)
	if err != nil {
		fmt.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}
	trx, err := r.UsersRepo.DB.Begin()
	defer trx.Rollback()
	if err != nil {
		fmt.Printf("error creating a transaction: %v", err)
		return nil, errors.New("something went wrong during DB Trx")
	}
	if _, err := r.UsersRepo.CreateUser(trx, user); err != nil {
		fmt.Printf("error creating a user: %v", err)
		return nil, err
	}
	if err = trx.Commit(); err != nil {
		fmt.Printf("error while committing trx: %v", err)
		return nil, err
	}
	token, err := user.GenereateToken()
	if err != nil {
		fmt.Printf("error while generating token: %v", err)
		return nil, errors.New("something went wrong")
	}
	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

//Login ...
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := r.UsersRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("email or password is invalid")
	}
	err = user.ComparePass(input.Password)
	if err != nil {
		return nil, errors.New("email or password is invalid")
	}
	token, err := user.GenereateToken()
	if err != nil {
		fmt.Printf("error while generating token: %v", err)
		return nil, errors.New("something went wrong")
	}
	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil

}
