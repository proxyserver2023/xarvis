package app

import (
	"github.com/sirupsen/logrus"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/access"
	"github.com/go-ozzo/ozzo-routing/fault"
	"github.com/go-ozzo/ozzo-validation"
	"net/http"

	"time"
)

func Init(logger *logrus.Logger) routing.Handler {
	return func(rc *routing.Context) error {
		now := time.Now()
		rc.Response = &access.LogResponseWriter{
			rc.Response,
			http.StatusOK,
			0,
		}
		ac := newRequestScope(now, logger, rc.Request)
		rc.Set("Context", ac)
		// fault.Recovery(ac.)
	}
}


// Get Response Scope returns the Request Scope of the current Request
func GetRequestScope(c *routing.Context) RequestScope {
	return c.Get("Context").(RequestScope)
}