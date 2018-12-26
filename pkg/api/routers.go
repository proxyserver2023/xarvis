package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", PingHandler)
	return router
}

type Response struct {
	Text string `json:"text"`
}

func respondJson(text string, w http.ResponseWriter) {
	response := Response{text}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	respondJson("Without Auth", w)
}

func SecuredPingHandler(w http.ResponseWriter, r *http.Request) {
	respondJson("Auth", w)
}
