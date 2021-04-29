package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping Request represents a request to run a command when /pin is called.
func Ping(c *gin.Context) {
	// return status 200(OK) and a String Pong
	c.JSON(http.StatusOK, "Pooong")
}
