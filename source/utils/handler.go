// File: `Server Handler Functions File` api/handlers.go
package PhoeniciaDigitalUtils

import (
	"encoding/json"
	"net/http"
)

type _PhoeniciaDigitalAPIFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Code       int `json:"statuss"`
	ErrorQuote any `json:"error"`
}

type ApiSuccess struct {
	Code         int `json:"status"`
	SuccessQuote any `json:"success"`
}

func SendJSON(w http.ResponseWriter, status int, val any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return json.NewEncoder(w).Encode(val)
}

func Handler(df _PhoeniciaDigitalAPIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		df(w, r)
	}
}
