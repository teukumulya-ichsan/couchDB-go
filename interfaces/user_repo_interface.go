package interfaces

import "github.com/teukumulya-ichsan/couchDB-go/models"

// UsersRepos interface ...
type UsersRepos interface {
	FindByID(docID string) (*models.User, error)
	SaveDocs(user models.User) (*models.User, error)
}
