package users

import (
	"strings"

	"github.com/FedericaCabrera1/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"data_created"`
	Status      string `json:"status"`
	Password    string `json:"password"` //dont want to be retreived when working with json
}
type Users []User

func (user *User) Validate() *errors.RestErr { //not a func but a method over the user, so the user validates itself and return if there´s an error or not
	user.FirstName = strings.TrimSpace(user.FirstName) //removing spaces
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
