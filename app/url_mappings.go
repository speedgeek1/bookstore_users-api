package app

import (
	"github.com/speedgeek1/bookstore_users-api/controllers/ping"
	"github.com/speedgeek1/bookstore_users-api/controllers/users"
)

func mapUrl() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:uid", users.GetUser)

	router.GET("/users/search", users.SearchUser)

	router.POST("/users", users.CreateUser)
}
