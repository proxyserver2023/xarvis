package apis

import (
	"github.com/alamin-mahamud/gapi/pkg/config"
)

// Start takes configuration struct and bootstraps the api
func Start(cfg *config.Configuration) {
	db, err := postgres.New(cfg.DB.PSN, cfg.DB.Timeout, cfg.DB.LogQueries)
	if err != nil {
		return err
	}

}
