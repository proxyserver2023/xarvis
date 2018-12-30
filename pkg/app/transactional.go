package app

import (
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/fault"
)

func Transactional(db *dbx.DB) routing.Handler {
	return func(c *routing.Context) error {
		tx, err := db.Begin()
		if err != nil {
			return err
		}

		rs := GetRequestScope(c)
		rs.SetTx(tx)

		err = fault.PanicHandler(rs.Errorf)(c)
		var e error
		if err != nil || rs.Rollback() {
			e = tx.Rollback()
		} else {
			e = tx.Commit()
		}

		if e != nil {
			if err == nil {
				return e
			}
			rs.Error(e)
		}

		return err
	}
}
