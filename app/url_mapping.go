package app

import (
	controller "github.com/Dhruv1993/bookstores_users_api/controllers"
)

// MapUrls is used to map all the route request to the controller functions
func MapUrls() {
	// you don't need to declare router from gin-gonic as it is already declared in application
	// which is the same package app.
	router.POST("/ping", controller.Ping)
	router.GET("/users/:user_id", controller.GetUser)
	// router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", controller.CreateUser)

}
