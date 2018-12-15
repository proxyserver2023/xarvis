package cmd

import (
	"fmt"

	"github.com/qiangxue/golang-restful-starter-kit/app"
)

// Run - Drives the whole app
func Run() {
	// 1. load app configuration
	if err := app.LoadConfig("../configs"); err != nil {
		panic(fmt.Errorf("Invalid Application Configuration"))
	}
	// 2. load error messages
	// 3. create the logger
	// 4. connect to the database
	// 5. wire up api routing
	// 6. start the Server
}
