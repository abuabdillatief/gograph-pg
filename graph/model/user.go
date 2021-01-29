package model

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" pg:",soft_delete"`
}

//HashPass ...
func (u *User) HashPass(password string) error {
	bytePass := []byte(password)
	hashed, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

//GenereateToken ...
func (u *User) GenereateToken() (*AuthToken, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 3) //3 days
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:       u.ID,
		IssuedAt: time.Now().Unix(),
		Issuer:   "meetmeup",
	})
	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	return &AuthToken{
		AccessToken: accessToken,
		ExpiredAt:   expiredAt,
	}, nil
}

//ComparePass ...
func (u *User) ComparePass(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

//HasRight ...
func (u *User) HasRight(m *Meetup) bool {
	return u.ID == m.UserID
}
