package interfaces

import "github.com/teukumulya-ichsan/couchDB-go/models"

type UsersRepos interface {
	FindById(docId string) (*models.User, error)
}
