package repository

import (
	"github.com/alamin-mahamud/xarvis/pkg/model"
	mgo "gopkg.in/mgo.v2"
)

// User - represents the base Repository
type JWT struct {
	C *mgo.Collection
}

func NewJWT(c *mgo.Collection) IJWT {
	return &JWT{
		c,
	}
}

func (j *JWT) Get() *model.User {
	return nil
}
