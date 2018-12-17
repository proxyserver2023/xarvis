package apis

import (
	"fmt"
	"net/http"
)

// Run runs the entire app
func Run() {
	fmt.Println("Server Running")
	http.ListenAndServe(":8081", nil)
}
