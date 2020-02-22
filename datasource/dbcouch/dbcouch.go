package dbcouch

import (
	"context"
	"log"

	"github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
)

// ConnectCouchDB init...
func ConnectCouchDB() *kivik.DB {
	client, err := kivik.New("couch", "http://localhost:5984/")
	if err != nil {
		log.Fatalf("Error Connect : %v", err)
	}

	err = client.Authenticate(context.Background(), couchdb.BasicAuth("root", "secret"))
	if err != nil {
		log.Fatalf("Error auth : %v", err)
	}

	db := client.DB(context.Background(), "fma")

	return db
}
