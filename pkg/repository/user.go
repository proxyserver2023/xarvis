package repository

import (
	"github.com/alamin-mahamud/xarvis/pkg/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User - represents the base Repository
type User struct {
	C *mgo.Collection
}

// NewUser - returns the implementation of IUser
func NewUser(c *mgo.Collection) IUser {
	return &User{c}
}

// List - lists all users
func (u *User) List() (*model.Users, error) {
	users := model.Users{}
	// err := u.C.Find(bson.M{}).All(users)

	iter := u.C.Find(nil).Iter()
	result := model.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}

	return &users, nil
}

// Create -
func (u *User) Create(m *model.User) error {
	objectID := bson.NewObjectId()
	m.ID = objectID
	err := u.C.Insert(m)
	return err

}

// Get -
func (u *User) Get(id string) (user *model.User, err error) {
	user = &model.User{}
	err = u.C.FindId(bson.ObjectIdHex(id)).One(user)
	return
}

// Update -
func (u *User) Update(id string, m *model.User) error {
	m.ID = bson.ObjectIdHex(id)
	err := u.C.Update(bson.M{"_id": bson.ObjectIdHex(id)}, m)
	return err
}

// Delete - user by id
func (u *User) Delete(id string) error {
	err := u.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
