package app

import (
	"github.com/FedericaCabrera1/bookstore_users-api/controllers/ping"
	"github.com/FedericaCabrera1/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping) //defining the function that need to be executed when we have a request of this path "/ping"
	//when receiving a request, we will apply the CONTROLLERS (we are defining which func inside of controllers package will be executed)
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get) /*meaning that when we have a GET request against users/ and a given user id,
	then the function responsible for handling this will be the GetUser func inside the controllers package */
	//	router.GET("/users/search", controllers.SearchUser)
	router.PUT("/users/:user_id", users.Update)    //PUT is to completely modify a given record
	router.PATCH("/users/:user_id", users.Update)  //PATCH is a partial modification, only what yo write, date, and everything will remain the same
	router.DELETE("/users/:user_id", users.Delete) //it will take that id and remove that user from the data base
	router.GET("/internal/users/search", users.Search)
}
