package app

import (
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type RequestScope interface {
	Logger
	UserID() string
	SetUserID(id string)
	RequestID() string
	Tx() *dbx.Tx
	SetTx(tx *dbx.Tx)
	Rollback() bool
	SetRollback(bool)
	Now() time.Time
}

type requestScope struct {
	Logger
	now       time.Time
	requestID string
	userID    string
	rollback  bool
	tx        *dbx.Tx
}

func (rs *requestScope) UserID() string {
	return rs.userID
}

func (rs *requestScope) SetUserID(id string) {
	rs.Logger.SetField("User ID", id)
	rs.userID = id
}

func (rs *requestScope) Tx() *dbx.Tx {
	return rs.tx
}

func (rs *requestScope) SetTx(tx *dbx.Tx) {
	rs.tx = tx
}

func (rs *requestScope) Rollback() bool {
	return rs.rollback
}

func (rs *requestScope) SetRollback(v bool) {
	rs.rollback = v
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func newRequestScope(now time.Time, logger *logrus.Logger, request *http.Request) RequestScope {
	l := NewLogger(logger, logrus.Fields{})
	requestID := request.Header.Get("X-Request-Id")
	if requestID != "" {
		l.SetField("RequestID", requestID)
	}
	return &requestScope{
		Logger:    l,
		now:       now,
		requestID: requestID,
	}
}
