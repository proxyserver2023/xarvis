package database

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

// NewDB - returns a mgo Collection
// In future this should be abstracted away with proper interfacing
// and must be dependency injected
func NewDB() *mgo.Collection {
	server := "localhost:27017"
	database := "xarvis-users"
	collection := "users"

	session, err := mgo.Dial(server)
	if err != nil {
		log.Fatal(err)
	}
	db := session.DB(database)
	c := db.C(collection)
	return c
}
