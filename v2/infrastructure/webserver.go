package infrastructure

import (
	"log"
	"net/http"
	"os"

	"github.com/alamin-mahamud/go-bootstrap/v2/infrastructure/router"
)

// Run - driver for the whole infrastructure
func Run() {
	r := router.Init()
	runWebServer(r)
}

func runWebServer(h http.Handler) {
	port := getPort()

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, h) // Goroutine will block here

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return port
}
