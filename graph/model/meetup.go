package model

//Meetup ...
type Meetup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
}

func (m *Meetup) HasRight(u *User) bool {
	return m.UserID == u.ID
}
