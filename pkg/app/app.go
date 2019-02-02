package app

import (
	"context"
	"log"
	"net/http"

	"github.com/alamin-mahamud/xarvis/pkg/database"
	"github.com/alamin-mahamud/xarvis/pkg/router"
)

// Run - bootstraps the whole app
func Run() {
	// config.Init()

	session, err := database.NewDB()
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
	}

}
