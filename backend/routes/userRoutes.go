package routes

import (
    "log"
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/wille1101/plant-waterer/backend/models"
    "github.com/wille1101/plant-waterer/backend/auth"
    db "github.com/wille1101/plant-waterer/backend/database"
)

// GenerateLoginToken handles the endpoint which logs a user in. If the provided credentials
// are correct a token is sent as a response.
func GenerateLoginToken(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    log.Println("Route accessed: login")

    var requestUser models.User

    d := json.NewDecoder(r.Body)
    d.DisallowUnknownFields()
    err := d.Decode(&requestUser)
    if (err != nil) {
        badRequestResp(w, r, err)
        return
    }

    user, err := db.GetUserQuery(requestUser)
    if (err != nil) {
        invalidLoginResp(w, r)
        return
    }

    checkPassword := user.CheckPassword(requestUser.Password)
    if (checkPassword != nil) {
        invalidLoginResp(w, r)
        return
    }

    tokenString, err := auth.GenerateJWT(user.UUID)
    if (err != nil) {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to generate JWT token")
        return
    }

    res := tokenResponse {
        Token: tokenString,
    }

    data, err := json.Marshal(res)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to encode JSON response")
        return
    }

    w.Write(data)
}

// RegisterUser handles the endpoint which registers a new user. A new user
// is created based on the credentials sent in the request.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    log.Println("Route accessed: registerUser")

    var user models.User

    d := json.NewDecoder(r.Body)
    d.DisallowUnknownFields()
    err := d.Decode(&user)
    if (err != nil) {
        badRequestResp(w, r, err)
        return
    }

    if (user.Password != user.ConfirmPassword) {
        passwordsNotMatchResp(w, r) 
        return
    }

    used := db.CheckUserNameQuery(user.UserName)
    if (used) {
        usernameTakenResp(w, r)
        return
    }

    err = user.HashPassword(user.Password)
    if (err != nil) {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to hash password")
        return
    }

    user.GenerateUUID()

    userId := db.InsertUserQuery(user)

    data, err := json.Marshal(userId)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Unable to encode JSON response")
        return
    }

    w.Write(data)
}
