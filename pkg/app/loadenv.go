package app

import (
	"os"

	"github.com/alamin-mahamud/gapi/pkg/daos"
)

func LoadEnvVars(dbConfig *daos.DBConfig) {
	dbConfig.DbUser = os.Getenv("DB_USER")
	dbConfig.DbName = os.Getenv("DB_NAME")
	dbConfig.DbPass = os.Getenv("DB_PASS")
	dbConfig.DbHost = os.Getenv("DB_HOST")
	dbConfig.DbPort = os.Getenv("DB_PORT")
}
