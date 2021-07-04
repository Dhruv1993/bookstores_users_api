package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSONString struct {
	Value string
	Valid bool
	Set   bool
}

// UnmarshalJSON helps in marshing json
func (i *JSONString) UnmarshalJSON(data []byte) error {
	fmt.Println("unmarshling is called")
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}
func (i *JSONString) Circle(data []byte) error {
	// The key isn't set to null
	fmt.Println("This is also run")
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

type Person struct {
	Email JSONString `json:"email"`
}

// Ping Request represents a request to run a command when /pin is called.
func Ping(c *gin.Context) {

	// fmt.Printf("%t the value is", person.Email.Set)
	// return status 200(OK) and a String Pong
	c.JSON(http.StatusOK, "pong")
}
