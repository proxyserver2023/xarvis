package usecase

import (
	"log"

	"github.com/alamin-mahamud/xarvis/pkg/model"
	"github.com/alamin-mahamud/xarvis/pkg/repository"
)

// User - represents the base Repository
type User struct {
	r repository.IUser
}

// NewUser - returns the implementation of IUser
func NewUser(r repository.IUser) IUser {
	return &User{r}
}

// List - lists all users
func (u *User) List() (*model.Users, error) {
	users, err := u.r.List()
	if err != nil {
		log.Fatalf("error %v", err)
	}
	return users, nil
}

// Create -
func (u *User) Create(m *model.User) error {
	err := u.r.Create(m)
	return err

}

// Get -
func (u *User) Get(id string) (user *model.User, err error) {
	user, err = u.r.Get(id)
	return user, err
}

// Update -
func (u *User) Update(id string, m *model.User) (err error) {
	err = u.r.Update(id, m)
	return err
}

// Delete - user by id
func (u *User) Delete(id string) (err error) {
	err = u.r.Delete(id)
	return err
}
