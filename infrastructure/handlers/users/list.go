package users

import (
	"fmt"
	"net/http"
)

// List - passes to the actual usecase for getting all the list with the repository.
func List(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Auth %s", "world")
}
