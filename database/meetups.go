package database

import (
	"fmt"

	"github.com/abuabdillatief/gograph-tutorial/graph/model"
	"github.com/go-pg/pg/v9"
)

//MeetupsRepo ...
type MeetupsRepo struct {
	DB *pg.DB
}

//GetMeetups ...
func (m *MeetupsRepo) GetMeetups(filter *model.MeetupFilter, limit, offset *int) ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	query := m.DB.Model(&meetups).Order("id")
	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query.Where("name ILIKE ?", fmt.Sprintf("%%%v%%", *filter.Name))
		}
	}
	if limit != nil {
		query.Limit(*limit)
	}
	if offset != nil {
		query.Offset(*offset)
	}
	err := query.Select()
	if err != nil {
		return nil, err
	}
	return meetups, err
}

//GetSelectedMeetups ...
func (m *MeetupsRepo) GetSelectedMeetups(id string) ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	err := m.DB.Model(&meetups).Where("user_id = ?", id).Order("id").Select()
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

//Delete ...
func (m *MeetupsRepo) Delete(id string) error {
	var meetup model.Meetup
	_, err := m.DB.Model(&meetup).Where("id= ?", id).Delete()
	fmt.Println(err)
	return err
}
