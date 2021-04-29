package controllers

import (
	"net/http"
	"strconv"

	"github.com/Dhruv1993/bookstores_users_api/domain/users"
	UserServices "github.com/Dhruv1993/bookstores_users_api/services"
	"github.com/Dhruv1993/bookstores_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateUser is used to create a user
func CreateUser(c *gin.Context) {
	var user users.User // this is a struct
	// This step gets the data that you post and does all the marshling and unmarshling
	// makes it equal to user data struct
	// This step is to compare the incoming data is equal to our user struct and assign the actual value to the
	// struct from the api response
	err := c.ShouldBindJSON(&user)
	if err != nil {
		//TODO
		restErr := errors.BadRequestError("json error")
		c.JSON(restErr.Status, restErr)
		return
	}

	// after unmarshalling and obtaining a clean parsed struct, create a user
	result, saveError := UserServices.CreateUser(user)

	if saveError != nil {
		//Todo handle the error
		c.JSON(saveError.Status, saveError)
		return
	}

	//fmt.Println(string(bytes))
	c.JSON(http.StatusCreated, result)
}

//GetUser is used to Get the user from end point
func GetUser(c *gin.Context) {
	userId, userError := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userError != nil {
		err := errors.BadRequestError("user_id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, err := UserServices.GetUser(userId)

	if err != nil {
		//Todo handle the error
		c.JSON(err.Status, err)
		return
	}

	//fmt.Println(string(bytes))
	c.JSON(http.StatusOK, user)

}

// SearchUser is used to find a user from end point
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement ME!!")

}
