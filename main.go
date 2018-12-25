package main

import (
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.Classic()
	n.Run(":8080")
}
