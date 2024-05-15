// File: `Server Handler Functions File` api/handlers.go
package PhoeniciaDigitalUtils

import (
	"encoding/json"
	"net/http"
)

type apifunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Code  int    `json:"statuss"`
	Error string `json:"error"`
}

type ApiSuccess struct {
	Code         int    `json:"status"`
	SuccessQuote string `json:"quote"`
}

func WriteJSON(w http.ResponseWriter, status int, val any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return json.NewEncoder(w).Encode(val)
}

func PhoeniciaDigitalHandler(df apifunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		df(w, r)
	}
}
