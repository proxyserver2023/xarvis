package repository

import "github.com/alamin-mahamud/xarvis/pkg/model"

// IJWT - represents the base Repository
type IJWT interface {
	Get() *model.User
}
