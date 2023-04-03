package middlewares

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/wille1101/plant-waterer/backend/auth"
)

// checkResponse is the structure of the response when an error occurs.
type checkResponse struct {
    ErrorType string `json:"errorType"`
	Message string `json:"message,omitempty"`
}

// CheckToken validates the token of an incoming request, making sure it's a valid token, before
// passing on the request to its handler function.
func CheckToken(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")

        token := r.Header.Get("Authorization")
        if (token == "") {
            res := checkResponse {
                ErrorType: "noToken",
                Message: "Header does not contain an Authorization token",
            }

            data, err := json.Marshal(res)
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                fmt.Fprintf(w, "Unable to encode JSON response")
                return
            }

            w.WriteHeader(http.StatusUnauthorized)
            w.Write(data)
            return
        }

        err := auth.ValidateToken(token)
        if (err != nil) {
            errorMsg := ""
            switch err.Error() {
            case "wrongAlg":
                errorMsg = "Unexpected signing method of the token."
            case "tokenExpired":
                errorMsg = "The JWT token has expired."
            case "invalidTokenLength":
                errorMsg = "The length of the token is invalid"
            case "invalidTokenChar":
                errorMsg = "The token contains an invalid character"
            default:
                errorMsg = err.Error()
            }

            res := checkResponse {
                ErrorType: err.Error(),
                Message: errorMsg,
            }

            data, err := json.Marshal(res)
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                fmt.Fprintf(w, "Unable to encode JSON response")
                return
            }

            w.WriteHeader(http.StatusUnauthorized)
            w.Write(data)
            return

        }

        next.ServeHTTP(w, r)
    })
}
