package app

import (
	"fmt"
)

func loadApplicationConfiguration() {
	if err := LoadConfig("config"); err != nil {
		panic(fmt.Errorf("Invalid Application Configuration"))
	}
}
