package apis

import (
	"fmt"
	"net/http"

	"github.com/alamin-mahamud/gapi/pkg/apis/routers"
	"github.com/alamin-mahamud/gapi/pkg/app"
	"github.com/alamin-mahamud/gapi/pkg/daos"
	"github.com/alamin-mahamud/gapi/pkg/utils"
)

func setDBConn() (daos.DB, error) {
	m := daos.NewMySQL()
	err := m.Open(dbConfig)
	utils.CheckErr(err)
	return m, err
}

var dbConfig = &daos.DBConfig{}

// Run runs the entire app
func Run() {
	app.LoadEnvVars(dbConfig)

	DB, err := setDBConn()
	utils.CheckErr(err)

	r := routers.BuildRouter(DB)
	fmt.Println("Server listening at 8081")
	http.ListenAndServe(":8081", r)
}
