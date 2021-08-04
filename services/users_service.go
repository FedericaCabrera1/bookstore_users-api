package services

import (
	"github.com/FedericaCabrera1/bookstore_users-api/domain/users"
	"github.com/FedericaCabrera1/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}    //a new instance of the user struct where result is a pointer to a new user with the id that we get as a parameter
	if err := result.Get(); err != nil { /*attempt to perform a get against the database, if theres an error, return nil and the error,
		if there´s no error, return the user and nil as an error */
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil { //validate that we dont have any error
		return nil, err
	}
	//at this point, the email is ready to be used on the database
	if err := user.Save(); err != nil { //attempt to save the user in the database
		return nil, err //if there´s an error, return nil and the error
	}
	return &user, nil //if we dont have any error, return a pointer to the user and nil
}
