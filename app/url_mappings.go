package app

import (
	"github.com/crjohnson1208/bookstore_users-api/controllers/ping"
	"github.com/crjohnson1208/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	//this get request is a direct query so we don't need to add :parameter
	router.GET("/internal/users/search", users.SearchUser)
}