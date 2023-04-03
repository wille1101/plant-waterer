package routes

import (
    "log"
    "fmt"
    "net/http"
    "encoding/json"
)

// tokenResponse is the response sent when a user logs in.
type tokenResponse struct {
	Token   string `json:"token"`
}

// response is a general response sent when, for example, a user is created.
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// userErrorResponse is the response sent when an error occurs when logging in or registering
// a new user.
type userErrorResponse struct {
    ErrorType string `json:"errorType"`
	Message string `json:"message,omitempty"`
}

// badRequestResp handles sending the reponse when the request is badly formatted.
func badRequestResp(w http.ResponseWriter, r *http.Request, err error) {
    res := userErrorResponse {
        ErrorType: "requestError",
        Message: "Unable to decode the body of the request: " + err.Error(),
    }

    data, err := json.Marshal(res)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to encode JSON response")
        return
    }

    w.WriteHeader(http.StatusBadRequest)
    w.Write(data)
}

// passwordsNotMatchResp handles sending the response when the password and confirm password in
// the request don't match.
func passwordsNotMatchResp(w http.ResponseWriter, r *http.Request) {
    res := userErrorResponse {
        ErrorType: "passwordsNotMatching",
        Message: "The password and confirm password don't match!",
    }

    data, err := json.Marshal(res)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to encode JSON response")
        return
    }

    w.WriteHeader(http.StatusUnauthorized)
    w.Write(data)
}

// invalidUsernameResp handles sending the response when the login information in the request is invalid.
func invalidLoginResp(w http.ResponseWriter, r *http.Request) {
    res := userErrorResponse {
        ErrorType: "invalidLogin",
        Message: "Invalid login information!",
    }

    data, err := json.Marshal(res)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to encode JSON response")
        return
    }

    w.WriteHeader(http.StatusUnauthorized)
    w.Write(data)
}

// usernameTakenResp handles sending the response when a user name provided in a register user request
// already is taken.
func usernameTakenResp(w http.ResponseWriter, r *http.Request) {
    res := userErrorResponse {
        ErrorType: "usernameTaken",
        Message: "The user name is already taken!",
    }

    data, err := json.Marshal(res)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to encode JSON response")
        return
    }

    w.WriteHeader(http.StatusConflict)
    w.Write(data)
}

// HomePage handles the base endpoint, /
func HomePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome!")
    log.Println("Route accessed: HomePage")
}

// setHeaders is a helper function which sets the required headers in the response.
func setHeaders(w http.ResponseWriter) (http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Headers", "Content-type, Authorization, Access-Control-Allow-Origin")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    return w
}

// OptionsGETOK sets the allowed methods header in the response to GET. Used for
// OPTIONS requests.
func OptionsGETOK(w http.ResponseWriter, r *http.Request) {
    w = setHeaders(w)
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.WriteHeader(http.StatusOK)
}

// OptionsPOSTOK sets the allowed methods header in the response to POST. Used for
// OPTIONS requests.
func OptionsPOSTOK(w http.ResponseWriter, r *http.Request) {
    w = setHeaders(w)
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.WriteHeader(http.StatusOK)
}

// OptionsDELETEOK sets the allowed methods header in the response to DELETE. Used for
// OPTIONS requests.
func OptionsDELETEOK(w http.ResponseWriter, r *http.Request) {
    w = setHeaders(w)
    w.Header().Set("Access-Control-Allow-Methods", "DELETE")
    w.WriteHeader(http.StatusOK)
}

// OptionsDELETEAndPOSTOK sets the allowed methods header in the response to DELETE and POST. Used for
// OPTIONS requests.
func OptionsDELETEAndPOSTOK(w http.ResponseWriter, r *http.Request) {
    w = setHeaders(w)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST")
    w.WriteHeader(http.StatusOK)
}

