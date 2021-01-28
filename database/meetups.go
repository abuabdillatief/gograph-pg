package database

import (
	"github.com/abuabdillatief/gograph-tutorial/graph/model"
	"github.com/go-pg/pg/v9"
)

type MeetupsRepo struct {
	DB *pg.DB
}

func (m *MeetupsRepo) GetMeetups() ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	err := m.DB.Model(&meetups).Select()
	if err != nil {
		return nil, err
	}
	return meetups, err
}

func (m *MeetupsRepo) CreateMeetup(obj *model.Meetup) (*model.Meetup, error) {
	_, err := m.DB.Model(obj).Returning("*").Insert()
	return obj, err
}
