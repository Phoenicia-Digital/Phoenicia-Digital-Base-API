// File: `Server Handler Functions File` source/utils/handlers.go
package PhoeniciaDigitalUtils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The Interface we use for our functions | Since they return an error
// used in http.Handle(_, <PhoeniciaDigitalHandler(function)>)
type PhoeniciaDigitalHandler func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	// ApiError struct can be returned as a error in API functions and automatically does error handling
	// for unsucessful calls || ERRORS via the ServeHTTP interface for `PhoeniciaDigitalHandler`
	Code       int `json:"status"`
	ErrorQuote any `json:"error"`
}

// Implements the `error` interface for ApiError.
func (apiErr ApiError) Error() string {
	if str, ok := apiErr.ErrorQuote.(string); ok {
		// Handle string ErrorQuote
		return fmt.Sprintf("Error Code: %d, Error Quote: %s", apiErr.Code, str)
	} else {
		// Handle non-string ErrorQuote
		return fmt.Sprintf("Error Code: %d, Error Quote (NON STRING VALUE): %v", apiErr.Code, apiErr.ErrorQuote)
	}
}

type ApiSuccess struct {
	// ApiSuccess struct can be returned as a `val` in `SendJSON` so we get a JSON response
	// for successful calls
	Code         int `json:"status"`
	SuccessQuote any `json:"success"`
}

func SendJSON(w http.ResponseWriter, status int, val any) error {
	// Set the Response Writers Header status and the content type to JSON so that we can send JSON
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)

	// Encode The Value `val` into the Response Writer and return an error if occured which will be managed by
	// the ServeHTTP interface ONLY IF USING `PhoeniciaDigitalHnadler`
	return json.NewEncoder(w).Encode(val)
}

// Implementaion of the http.ServeHTTP interface on `PhoeniciaDigitalHandler`
func (pdf PhoeniciaDigitalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Call the underlying handler function & Handle Errors
	if err := pdf(w, r); err != nil {
		Log(err.Error())
		if apiErr, ok := err.(ApiError); ok {
			if ierr := SendJSON(w, apiErr.Code, apiErr); ierr != nil {
				http.Error(w, apiErr.Error(), apiErr.Code)
			}
		} else {
			if ierr := SendJSON(w, http.StatusInternalServerError, ApiError{Code: http.StatusInternalServerError, ErrorQuote: err.Error()}); ierr != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

//``` DISCONTINUED USED TO USE THIS AS THE HANDLER INSIDE THE http.HandleFunc but now we http.Handle with the
// PhoeniciaDigitalAPIFunction Interface! ```

// func Handler(df _PhoeniciaDigitalAPIFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
// 		df(w, r)
// 	}
// }
