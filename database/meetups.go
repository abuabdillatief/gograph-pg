package database

import (
	"github.com/abuabdillatief/gograph-tutorial/graph/model"
	"github.com/go-pg/pg/v9"
)

//MeetupsRepo ...
type MeetupsRepo struct {
	DB *pg.DB
}

//GetMeetups ...
func (m *MeetupsRepo) GetMeetups() ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	err := m.DB.Model(&meetups).Order("id").Select()
	if err != nil {
		return nil, err
	}
	return meetups, err
}

//GetByID ...
func (m *MeetupsRepo) GetByID(id string) (*model.Meetup, error) {
	var meetup model.Meetup
	err := m.DB.Model(&meetup).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &meetup, err
}

//CreateMeetup ...
func (m *MeetupsRepo) CreateMeetup(obj *model.Meetup) (*model.Meetup, error) {
	_, err := m.DB.Model(obj).Returning("*").Insert()
	return obj, err
}

//Update ...
func (m *MeetupsRepo) Update(obj *model.Meetup) (*model.Meetup, error) {
	_, err := m.DB.Model(obj).Where("id= ?", obj.ID).Update()
	return obj, err
}
