package app

import (
	"github.com/crjohnson1208/bookstore_users-api/controllers/ping"
	"github.com/crjohnson1208/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", users.SearchUser)
	router.POST("/users" controllers.CreateUser)
}