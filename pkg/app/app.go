package app

import (
	"flag"

	"github.com/alamin-mahamud/gapi/pkg/api"
	"github.com/alamin-mahamud/gapi/pkg/config"
)

// Run Drives the Entire Application
func Run() error {
	cfgPath := flag.String(
		"p",
		"./config/local.yaml",
		"Path to config file",
	)

	flag.Parse()
	cfg, err := config.Load(*cfgPath)
	checkErr(err)
	checkErr(api.Start())
}
