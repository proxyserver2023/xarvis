package repository

import (
	"github.com/alamin-mahamud/xarvis/pkg/model"
	mgo "gopkg.in/mgo.v2"
)

// User - represents the base Repository
type User struct {
	C *mgo.Collection
}

// NewUser - returns the implementation of IUser
func NewUser(c *mgo.Collection) IUser {
	return &User{
		C: c,
	}
}

func (u *User) List() []*model.User {}
func (u *User) Create() *model.User {}
func (u *User) Get() *model.User    {}
func (u *User) Update() *model.User {}
func (u *User) Delete() *model.User {}
