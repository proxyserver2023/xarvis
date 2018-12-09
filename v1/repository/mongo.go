package repository

import (
	mgo "gopkg.in/mgo.v2"

	articleEntity "github.com/alamin-mahamud/go-bootstrap/v1/entity/article"
)

// MongoDB holds the MongoDB Session
type MongoDB struct {
	Db *mgo.Session
}

// Open -> Open the MongoDB Session
func (m *MongoDB) Open() {
	Db, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	m.Db = Db
	m.Db.SetMode(mgo.Monotonic, true)
}

// Close -> Close the MongoDB session
func (m *MongoDB) Close() {
	m.Db.Close()
}

// List
func (m *MongoDB) List() ([]articleEntity.Article, error) {
	articleList := []articleEntity.Article{
		{ID: 1, Name: "first article"},
		{ID: 2, Name: "second article"},
		{ID: 3, Name: "third article"},
	}
	return articleList, nil
}
