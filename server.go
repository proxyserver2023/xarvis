package gapi

import (
	"fmt"
	"net/http"

	"github.com/alamin-mahamud/gapi/pkg/app"
	"github.com/alamin-mahamud/gapi/errors"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/sirupsen/logrus"
)

func main() {
	// load application configuration
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid Applciation Configuration: %s", err))
	}

	// load error messages
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// create the logger
	logger := logrus.New()

	// connect to the database
	db, err := dbx.MustOpen("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}
	db.LogFunc = logger.Infof

	// wire up API Routing
	http.Handle("/", buildRouter(logger, db))

	// start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("Server %v is started at  %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}
