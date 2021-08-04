package users

import (
	"strings"

	"github.com/FedericaCabrera1/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"data_created"`
}

func (user *User) Validate() *errors.RestErr { //not a func but a method over the user, so the user validates itself and return if there´s an error or not
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
