package repositories

import (
	"context"
	"github.com/go-kivik/kivik"
	"github.com/teukumulya-ichsan/couchDB-go/datasource/couch_db"
	"github.com/teukumulya-ichsan/couchDB-go/models"
)

type CouchRepo struct {
	db *kivik.DB
}

func NewCouchRepo() *CouchRepo {
	return &CouchRepo{
		couch_db.ConnectCouchDB(),
	}
}

func (c *CouchRepo) FindById(docId string) (*models.User, error) {
	row := c.db.Get(context.Background(), docId)
	_ = c.db.Close(context.Background())

	var dataUser *models.User
	if err := row.ScanDoc(&dataUser); err != nil {
		return nil, err
	}

	return dataUser, nil
}
