package services

import (
	UsersStruct "github.com/Dhruv1993/bookstores_users_api/domain/users"
	"github.com/Dhruv1993/bookstores_users_api/utils/errors"
)

func GetUser(userId int64) (*UsersStruct.User, *errors.RestError) {
	// make a user struct with only userId
	result := &UsersStruct.User{Id: userId}
	// call GetUser method in user_dao and create a filled user object with results from DB
	// when you call this, you end up filling user object struct with the info from DB
	if err := result.GetUser(); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser receive a user which is of type User and return pointer to the interface and error, store the address as &user in the *UsersStruct.User
func CreateUser(user UsersStruct.User) (*UsersStruct.User, *errors.RestError) {
	// validate the user first
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	// save the user to the DB if there are no errors
	if err := user.SaveUser(); err != nil {
		return nil, err
	}
	return &user, nil
}

//Live Demo
//#include <stdio.h>
//
//int main () {

//int  var = 20;   /* actual variable declaration */
//int  *ip;        /* pointer variable declaration */
//
//ip = &var;  /* store address of var in pointer variable*/
//
//printf("Address of var variable: %x\n", &var  );
//
///* address stored in pointer variable */
//printf("Address stored in ip variable: %x\n", ip );
//
///* access the value using the pointer */
//printf("Value of *ip variable: %d\n", *ip );
//
//return 0;
//}
