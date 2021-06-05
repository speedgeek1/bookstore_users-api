package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/speedgeek1/bookstore_users-api/domain/users"
	"github.com/speedgeek1/bookstore_users-api/services"
	"github.com/speedgeek1/bookstore_users-api/utills/errors"
)

func GetUser(c *gin.Context) {
	userId, usrErr := strconv.ParseInt(c.Param("uid"), 10, 64)
	if usrErr != nil {
		err := errors.NewBadRequestError("Invalid user Id! UserId should be an integer")
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

func CreateUser(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		// TODO: handle json error
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO : Handle database error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusAccepted, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
