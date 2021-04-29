// dao is data access object
// this part of the code is purely for data persistence that is connecting to a DB and retreving a user
package users

import (
	"fmt"

	"github.com/Dhruv1993/bookstores_users_api/utils/errors"
)

// mimic a DB for now
// this is like {1:{first_name}}
var usersDB = make(map[int64]*User)

// sinc we want GetUser to work with user object and should also have the power to change
// its keys and values we pass a pointer to the User. This will populate the user object
func (user *User) GetUser() *errors.RestError {
	// .. find the user by Id
	result := usersDB[user.Id]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("no user with id %d found", user.Id))
	}
	// change the structure of the user and copy the content that we found within the DB
	user.Id = result.Id
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated

	return nil
}

// save the user to the DB
func (user *User) SaveUser() *errors.RestError {
	// check if the user exists
	current := usersDB[user.Id]
	// if current user object is not nil or it means it is existing is the db
	if current != nil {
		if current.Email == user.Email {
			return errors.BadRequestError(fmt.Sprintf("Submitted Email: %s is already registered", user.Email))
		}
		return errors.BadRequestError(fmt.Sprintf("user %d already exixxts", user.Id))
	}

	usersDB[user.Id] = user

	for key, element := range usersDB {
		fmt.Println("key", key, "=>", "Element", element)
	}

	return nil
}
