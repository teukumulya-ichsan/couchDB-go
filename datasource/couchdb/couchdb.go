package couchdb

import (
	"context"
	"github.com/go-kivik/couchdb"
	_ "github.com/go-kivik/couchdb" // The CouchDB driver
	"github.com/go-kivik/kivik"
	"log"
)

func ConnectCouchDB() *kivik.DB {
	client, err := kivik.New("couch", "http://localhost:5984/")
	if err != nil {
		log.Fatalf("Error Connect : %v", err)
	}

	err = client.Authenticate(context.Background(), couchdb.BasicAuth("root", "secret"))
	//db := client.DB(context.Background(), "db1")
	//err = db.Client().Authenticate(context.Background(), couchdb.BasicAuth("root", "secret"))
	if err != nil {
		log.Fatalf("Error auth : %v", err)
	}

	db :=  client.DB(context.Background(), "db1")

	return db
}
