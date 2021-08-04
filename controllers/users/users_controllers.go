package users

import (
	"net/http"

	"strconv"

	"github.com/FedericaCabrera1/bookstore_users-api/domain/users"
	"github.com/FedericaCabrera1/bookstore_users-api/services"
	"github.com/FedericaCabrera1/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) { /*we add the argument c being a pointer of *gin.Context to add the interface,
	in order to use the func as type Handlerfunc*/
	var user users.User                             //given user that is part of the domain
	if err := c.ShouldBindJSON(&user); err != nil { //jason request to bind the info received into the user struct
		restErr := errors.NewBadRequestError("invalid json body") /*if we have errors we return the
		badrequest struct with that message*/
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user) /*if there´s no error,
	we use the user and send it to a service in order to save it */
	if saveErr != nil { //if there are errors when saving it we return a jason w the error
		c.JSON(saveErr.Status, saveErr)
		//Handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result) /*if we don´t have any errors,
	we just return an status created with the result (the result when creating the user, is a pointer to user)*/
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64) //if there is an error trying to parse the user id parameter that we are receiving, then we return a bad request
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

/*
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}*/
