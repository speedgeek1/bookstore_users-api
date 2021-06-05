package users

import (
	"fmt"

	"github.com/speedgeek1/bookstore_users-api/utills/date_utills"
	"github.com/speedgeek1/bookstore_users-api/utills/errors"
)

var userDb = make(map[int64]*User)

func (user *User) Save() *errors.RestErr {

	current := userDb[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}

	user.DateCreated = date_utills.GetNowString()
	userDb[user.Id] = user
	return nil
}

func (user *User) Get() *errors.RestErr {
	result := userDb[user.Id]
	if result == nil {
		return errors.NotFoundError("User not found")
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil

}
