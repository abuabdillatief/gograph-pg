package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/abuabdillatief/gograph-tutorial/graph/model"
)

//Register ...
func (d *Domain) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	_, err := d.UsersRepo.GetUserByEmail(input.Email)
	//we are expecting errors, as it shows that user doesn't exist
	if err == nil {
		return nil, errors.New("email is already in used")
	}
	_, err = d.UsersRepo.GetUserByUsername(input.Username)
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
	trx, err := d.UsersRepo.DB.Begin()
	defer trx.Rollback()
	if err != nil {
		fmt.Printf("error creating a transaction: %v", err)
		return nil, errors.New("something went wrong during DB Trx")
	}
	if _, err := d.UsersRepo.CreateUser(trx, user); err != nil {
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
func (d *Domain) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := d.UsersRepo.GetUserByEmail(input.Email)
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
