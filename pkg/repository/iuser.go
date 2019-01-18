package repository

import "github.com/alamin-mahamud/xarvis/pkg/model"

// IUser - represents the base Repository
type IUser interface {
	List() []*model.User
	Create() *model.User
	Get() *model.User
	Update() *model.User
	Delete() *model.User
}
