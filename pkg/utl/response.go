package utl

import (
	"encoding/json"
	"net/http"
)

// SendJSONResponse - ...
func SendJSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// SendJSONErrResponse - ...
func SendJSONErrResponse(w http.ResponseWriter, code int, msg string) {
	SendJSONResponse(w, code, map[string]string{"error": msg})
}
