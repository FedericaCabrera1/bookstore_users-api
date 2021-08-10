package services

import (
	"github.com/FedericaCabrera1/bookstore_users-api/domain/users"
	"github.com/FedericaCabrera1/bookstore_users-api/utils/crypto_utils"
	"github.com/FedericaCabrera1/bookstore_users-api/utils/date_utils"
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
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	//at this point, the email is ready to be used on the database
	if err := user.Save(); err != nil { //attempt to save the user in the database
		return nil, err //if there´s an error, return nil and the error
	}
	return &user, nil //if we dont have any error, return a pointer to the user and nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id) //before updating the user, attempt to get that user from the data base
	if err != nil {
		return nil, err
	}
	//if I reach this point I am going to use the info inside the user we get as a parameter and attempt to complete the current one and save it
	//is this a partial request? if so then firs
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
