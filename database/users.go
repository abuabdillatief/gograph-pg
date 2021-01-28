package database

import (
	"github.com/abuabdillatief/gograph-tutorial/graph/model"
	"github.com/go-pg/pg/v9"
)

//UsersRepo ...
type UsersRepo struct {
	DB *pg.DB
}

//GerUserByID ...
func (u *UsersRepo) GerUserByID(id string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
