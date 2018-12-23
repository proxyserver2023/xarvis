package main

import (
	"runtime"

	"github.com/alamin-mahamud/gapi/pkg/app"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app.Bootstrap()
}
