package repositories

import (
	"context"

	"github.com/go-kivik/kivik"
	"github.com/teukumulya-ichsan/couchDB-go/datasource/dbcouch"
	"github.com/teukumulya-ichsan/couchDB-go/models"
)

// CouchRepo struct ...
type CouchRepo struct {
	db *kivik.DB
}

// NewCouchRepo constructor ...
func NewCouchRepo() *CouchRepo {
	return &CouchRepo{
		dbcouch.ConnectCouchDB(),
	}
}

// FindByID method (mutation) ...
func (c *CouchRepo) FindByID(docID string) (*models.User, error) {
	row := c.db.Get(context.Background(), docID)
	_ = c.db.Close(context.Background())

	var dataUser *models.User
	if err := row.ScanDoc(&dataUser); err != nil {
		return nil, err
	}

	return dataUser, nil
}
