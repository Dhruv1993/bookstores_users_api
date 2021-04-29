package users

import (
	"strings"

	"github.com/Dhruv1993/bookstores_users_api/utils/errors"
	RestErrors "github.com/Dhruv1993/bookstores_users_api/utils/errors"
)

// User what it means is use the id key from the json that we want to parse and populate the Id field of the struct
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//Validate the user assign struct to a method Validate
// this is a type of pointer receiver and can modify the actual struct
// There are two type of receiver for the structs - value and pointer
//if this (user *User) is passed as (user User) it becomes value receiver and
//you can't change the actual struct
// this is like a struct has a method defined for it that can be accessed via package user

/*
Go does not have classes. However, you can define methods on types.
A method is a function with a special receiver argument.
The receiver appears in its own argument list between the func keyword and the method name.
In this example, the Abs method has a receiver of type Vertex named v.
*/

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return RestErrors.BadRequestError("Invalid email")
	}
	return nil
}
