package repository

import "github.com/alamin-mahamud/xarvis/pkg/model"

// IUser - represents the base Repository
type IUser interface {
	List() (*model.Users, error)
	Create(*model.User) error
	Get(string) (*model.User, error)
	Update(string, *model.User) error
	Delete(string) error
}
