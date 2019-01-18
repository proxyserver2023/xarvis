package app

import (
	"log"
	"net/http"

	"github.com/alamin-mahamud/xarvis/pkg/router"
)

func Run() {
	// config.Init()
	router := router.Init()
	http.Handle("/", router)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
