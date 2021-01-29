package domain

import (
	"github.com/abuabdillatief/gograph-tutorial/database"
	"github.com/abuabdillatief/gograph-tutorial/graph/model"
)

//Domain ...
type Domain struct {
	UsersRepo   database.UsersRepo
	MeetupsRepo database.MeetupsRepo
}

//NewDomain ...
//this is called a constructor function, a func to create a new instance
func NewDomain(usersRepo database.UsersRepo, meetupsRepo database.MeetupsRepo) *Domain {
	return &Domain{
		UsersRepo:   usersRepo,
		MeetupsRepo: meetupsRepo,
	}
}

//Owner ...
type Owner interface {
	HasRight(user *model.User) bool
}

func checkOwnership(o Owner, u *model.User) bool {
	return o.HasRight(u)
}
