package database

import (
	mgo "gopkg.in/mgo.v2"
)

const (
	server     = "localhost:27017"
	database   = "xarvis-users"
	collection = "users"
)

// NewDB - returns a mgo Session
// In future this should be abstracted away with proper interfacing
// and must be dependency injected
func NewDB() (*mgo.Session, error) {
	session, err := mgo.Dial(server)
	return session, err
}

/*
db := session.DB(database)
	c := db.C(collection)
	return c
*/
