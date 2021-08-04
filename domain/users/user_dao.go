package users

import (
	"fmt"

	"github.com/FedericaCabrera1/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr { //get a user from the database based on the user id
	//work with a pointer in order to make sure we are modifing the actual element and not the copy
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError((fmt.Sprintf("user %d not found", user.Id)))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr { //as we are in the domain package, the user knows how to save itself, so we create a method, not a func
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	usersDB[user.Id] = user //if there are not any errors, we save the user in the database and return nil
	return nil
}
