package gapi

import (
	"fmt"
	"github.com/qiangxue/golang-restful-starter-kit/app"
)

func main() {
	// load application configuration
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid Applciation Configuration: %s", err))
	}
}