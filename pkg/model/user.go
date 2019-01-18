package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	UserName string        `json:"username"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
}

type Users []User
