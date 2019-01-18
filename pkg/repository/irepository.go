package repository

import "github.com/alamin-mahamud/xarvis/pkg/model"

// Base
type IRepository interface {
	List() []*model.User
	Create() *model.User
	Get() *model.User
	Update() *model.User
	Delete() *model.User
}
