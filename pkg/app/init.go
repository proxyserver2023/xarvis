package app

import (
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/access"
	"github.com/go-ozzo/ozzo-routing/fault"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/sirupsen/logrus"
	"database/sql"
	"fmt"
	"net/http"

	"time"
	"github.com/alamin-mahamud/gapi/errors"
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
		fault.Recovery(ac.Errorf, convertError)(rc)
		logAccess(rc, ac.Infof, ac.Now())
		return nil
	}
}

// Get Response Scope returns the Request Scope of the current Request
func GetRequestScope(c *routing.Context) RequestScope {
	return c.Get("Context").(RequestScope)
}
func convertError(c *routing.Context, err error) error {
	if err == sql.ErrNoRows {
		return errors.NotFound("The Requested resource")
	}

	switch err.(type) {
	case *errors.APIError:
		return err
	case validation.Errors:
		return errors.InvalidData(err.(validation.Errors))
	case routing.HTTPError:
		switch err.(routing.HTTPError).StatusCode() {
		case http.StatusUnauthorized:
			return errors.Unauthorized(err.Error())
		case http.StatusNotFound:
			return errors.NotFound("The requested resource")
		}
	}
	return errors.InternalServerError(err)
}
func logAccess(c *routing.Context, logFunc access.LogFunc, start time.Time) {
	rw := c.Response.(*access.LogResponseWriter)
	elapsed := float64(time.Now().Sub(start).Nanoseconds())/1e6
	requestLine := fmt.Sprintf(
		"%s %s %s",
		c.Request.Method,
		c.Request.URL.Path,
		c.Request.Proto,
		)
	logFunc(`[%.3fms] %s %d %d`, elapsed, requestLine, rw.Status, rw.BytesWritten)
}