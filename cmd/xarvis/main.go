package main

import "github.com/alamin-mahamud/xarvis/pkg/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
