package handler

import (
	"net/http"

	"github.com/alamin-mahamud/xarvis/pkg/usecase"
)

type user struct {
	u usecase.IUser
}

// NewUser - returns IHandler implementation of user
func NewUser(uu usecase.IUser) IHandler {
	return &user{
		u: uu,
	}
}

// List - /v1/authentication/users
func (u *user) List(w http.ResponseWriter, r *http.Request) {

}

// Create - /v1/authentication/users
func (u *user) Create(w http.ResponseWriter, r *http.Request) {}

// Get - /v1/authentication/users/<id>
func (u *user) Get(w http.ResponseWriter, r *http.Request) {}

// Update - /v1/authentication/users/<id>
func (u *user) Update(w http.ResponseWriter, r *http.Request) {}

// Delete - /v1/authentication/users/<id>
func (u *user) Delete(w http.ResponseWriter, r *http.Request) {}
