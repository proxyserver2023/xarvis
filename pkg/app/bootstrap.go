package app

import (
	"log"
	"net/http"
	"os"
)

// Bootstrap the api
func Bootstrap() {
	port := os.Getenv("APP_PORT")
	http.Handle("/", api.Handlers())
	log.Printf("Server up on Port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
