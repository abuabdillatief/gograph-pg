package database

import (
	"fmt"

	"github.com/abuabdillatief/gograph-tutorial/graph/model"
	"github.com/go-pg/pg/v9"
)

//UsersRepo ...
type UsersRepo struct {
	DB *pg.DB
}

//GetUserByField ...
func (u *UsersRepo) GetUserByField(field, value string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &user, err
}

//GetUserByID ...
func (u *UsersRepo) GetUserByID(id string) (*model.User, error) {
	return u.GetUserByField("id", id)
}

//GetUserByEmail ...
func (u *UsersRepo) GetUserByEmail(email string) (*model.User, error) {
	return u.GetUserByField("email", email)
}

//GetUserByUsername ...
func (u *UsersRepo) GetUserByUsername(username string) (*model.User, error) {
	return u.GetUserByField("username", username)
}

//CreateUser ...
func (u *UsersRepo) CreateUser(trx *pg.Tx, user *model.User) (*model.User, error) {
	_, err := trx.Model(user).Returning("*").Insert()
	return user, err
}
