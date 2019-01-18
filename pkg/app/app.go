package app

import (
	"log"
	"net/http"

	"github.com/alamin-mahamud/xarvis/pkg/router"
)

// Run - bootstraps the whole app
func Run() {
	// config.Init()
	router := router.Init()
	http.Handle("/", router)
	log.Println("Starting the Server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
