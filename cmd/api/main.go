package main

import (
	"github.com/alamin-mahamud/gapi/pkg/app"
)

func main() {
	checkErr(app.Run())
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
