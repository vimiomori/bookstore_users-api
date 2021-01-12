package app

import (
	"github.com/vimiomori/bookstore_users-api/controllers/ping"
	"github.com/vimiomori/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users:user_id", users.CreateUser)
	// router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}
