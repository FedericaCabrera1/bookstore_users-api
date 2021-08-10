package users

import (
	"encoding/json"
)

type PublicUser struct {
	//we just want to retreive id, date and status
	Id int64 `json:"id"`
	//FirstName   string `json:"first_name"`  (hidden)
	//LastName    string `json:"last_name"`  (hidden)
	//Email       string `json:"email"`  (hidden)
	DateCreated string `json:"data_created"`
	Status      string `json:"status"`
	//Password    string `json:"password"` (password is hidden in both cases)
}

type PrivateUser struct {
	//we  want to retreive everything but the password
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"data_created"`
	Status      string `json:"status"`
	//Password    string `json:"password"` (hidden)
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser

}
