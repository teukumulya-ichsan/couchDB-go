package app

import "github.com/teukumulya-ichsan/couchDB-go/controllers/users"

var (
	userController = users.NewUserController()
)

func mapsUrls() {
	httpRouter.GET("/user/{id}", userController.GetUserById)
}
