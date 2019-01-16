package mock

import (
	"github.com/alamin-mahamud/gapi/pkg/utl/model"
)

// JWT mock
type JWT struct {
	GenerateTokenFn func(*model.User) (string, string, error)
}

// GenerateToken mock
func (j *JWT) GenerateToken(u *model.User) (string, string, error) {
	j.GenerateTokenFn(u)
}
