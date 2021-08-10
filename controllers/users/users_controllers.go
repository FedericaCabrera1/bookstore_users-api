package users

import (
	"net/http"

	"strconv"

	"github.com/FedericaCabrera1/bookstore_users-api/domain/users"
	"github.com/FedericaCabrera1/bookstore_users-api/services"
	"github.com/FedericaCabrera1/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	//receives an id as a parameter
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64) //attempts to parse it, if theres an error, return a bad request error
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil //if there´s no error, we return the user id and nil
}

func Create(c *gin.Context) { /*we add the argument c being a pointer of *gin.Context to add the interface,
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

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true")) /*if we don´t have any errors,
	we just return an status created with the result (the result when creating the user, is a pointer to user)*/
}

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id")) //if there is an error trying to parse the user id parameter that we are receiving, then we return a bad request
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id")) //if there is an error trying to parse the user id parameter that we are receiving, then we return a bad request
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User                             //given user that is part of the domain
	if err := c.ShouldBindJSON(&user); err != nil { //jason request to bind the info received into the user struct
		restErr := errors.NewBadRequestError("invalid json body") /*if we have errors we return the
		badrequest struct with that message*/
		c.JSON(restErr.Status, restErr)
		return
	}
	//if we reach this point, we have the user with all the fields populated
	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch   //is partial is going to be true in the case we are sending a PATCH request instead of a PUT
	result, err := services.UpdateUser(isPartial, user) //attempt to update the user
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
