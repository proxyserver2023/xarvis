package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alamin-mahamud/xarvis/pkg/model"
	"github.com/alamin-mahamud/xarvis/pkg/usecase"
	"github.com/alamin-mahamud/xarvis/pkg/utl"
	"github.com/gorilla/mux"
)

type user struct {
	service usecase.IUser
}

// NewUser - returns IHandler implementation of user
func NewUser(uu usecase.IUser) IHandler {
	return &user{
		service: uu,
	}
}

// List - /v1/authentication/users
func (u *user) List(w http.ResponseWriter, r *http.Request) {
	users, err := u.service.List()

	if err != nil {
		utl.SendJSONErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utl.SendJSONResponse(w, http.StatusOK, *users)
}

// Create - /v1/authentication/users
func (u *user) Create(w http.ResponseWriter, r *http.Request) {
	// defer r.Body.Close() - No need to use it cause server close the request by itself

	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utl.SendJSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := u.service.Create(&user); err != nil {
		utl.SendJSONErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utl.SendJSONResponse(w, http.StatusCreated, user)
}

// Get - /v1/authentication/users/<id>
func (u *user) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		utl.SendJSONErrResponse(w, http.StatusBadRequest, "You must Provide an ID")
		return
	}

	user, err := u.service.Get(id)
	if err != nil {
		utl.SendJSONErrResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utl.SendJSONResponse(w, http.StatusOK, *user)
}

// Update - /v1/authentication/users/<id>
func (u *user) Update(w http.ResponseWriter, r *http.Request) {
	// defer r.Body.Close() - No need to use it cause server close the request by itself
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		utl.SendJSONErrResponse(w, http.StatusBadRequest, "You must Provide an ID")
		return
	}

	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utl.SendJSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := u.service.Update(id, &user); err != nil {
		utl.SendJSONErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utl.SendJSONResponse(w, http.StatusOK, map[string]string{"result": "success"})
}

// Delete - /v1/authentication/users/<id>
func (u *user) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		utl.SendJSONErrResponse(w, http.StatusBadRequest, "You must Provide an ID")
		return
	}

	if err := u.service.Delete(id); err != nil {
		utl.SendJSONErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utl.SendJSONResponse(w, http.StatusOK, map[string]string{"result": "Successfully Deleted"})
}
