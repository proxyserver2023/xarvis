package app

import (
	"fmt"
	"net/http"

	"github.com/alamin-mahamud/xarvis/pkg/errors"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	_ "github.com/lib/pq" // postgresql native driver
	"github.com/sirupsen/logrus"
)

// Run - bootstraps the whole app
func Run() error {
	// load application configurations
	if err := loadConfig("config"); err != nil {
		return fmt.Errorf("Invalid application configuration: %s", err)
	}

	// load  error messages
	if err := errors.LoadMessages(Config.ErrorFile); err != nil {
		return fmt.Errorf("Failed to read the error message file: %s", err)
	}

	// connect to the database
	db, err := dbx.MustOpen("postgres", Config.DSN)
	if err != nil {
		panic(err)
	}

	// create the logger
	logger := logrus.New()
	db.LogFunc = logger.Infof

	// wire up API routing
	http.Handle("/", buildRouter(logger, db))

	/* session, err := database.NewDB()
	defer session.Close()

	mPool := mgoSession.NewPool(context.Background(), session, 10)
	defer mPool.Close()

	br := bookmarkRepository.NewMongoRepository(mPool, config.MONGODB_DATABASE)
	bu := bookmarkUsecase.NewService(br)

	router := router.Init()
	http.Handle("/", router)
	log.Println("Starting the Server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	} */

	return nil
}

func buildRouter(logger *logrus.Logger, db *dbx.DB) *routing.Router {
	router := routing.New()

	router.Use(
		InitScope(logger),
		content.TypeNegotiator(content.JSON),
	)

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		c.Abort() // skip all other middlewares/handlers
		return c.Write("OK " + Version)
	})

	return router
}
