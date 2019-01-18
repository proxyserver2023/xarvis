package handler

import (
	"fmt"
	"net/http"
)

type user struct{}

// NewUser - returns IHandler implementation of user
func NewUser() IHandler {
	return &user{}
}

// List - /v1/authentication/users
func (u *user) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello in Web")
}

// Create - /v1/authentication/users
func (u *user) Create(w http.ResponseWriter, r *http.Request) {}

// Get - /v1/authentication/users/<id>
func (u *user) Get(w http.ResponseWriter, r *http.Request) {}

// Update - /v1/authentication/users/<id>
func (u *user) Update(w http.ResponseWriter, r *http.Request) {}

// Delete - /v1/authentication/users/<id>
func (u *user) Delete(w http.ResponseWriter, r *http.Request) {}
