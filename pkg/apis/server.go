package apis

import (
	"fmt"
	"net/http"
)

func Run() {
	fmt.Println("Server Running")
	http.ListenAndServe(":8081", nil)
}
