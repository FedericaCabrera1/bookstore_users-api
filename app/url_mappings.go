package app

import (
	"github.com/FedericaCabrera1/bookstore_users-api/controllers/ping"
	"github.com/FedericaCabrera1/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping) //defining the function that need to be executed when we have a request of this path "/ping"
	//when receiving a request, we will apply the CONTROLLERS (we are defining which func inside of controllers package will be executed)
	router.GET("/users/:user_id", users.GetUser) /*meaning that when we have a GET request against users/ and a given user id,
	then the function responsible for handling this will be the GetUser func inside the controllers package */
	//	router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}
