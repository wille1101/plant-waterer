package database

import (
    "log"

    "github.com/wille1101/plant-waterer/backend/models"

    "github.com/google/uuid"
)

// GetUserQuery queries the database for a users information given a user name.
func GetUserQuery(requestUser models.User) (models.User, error) {
    db := createConnection()
    defer db.Close()

    sqlQuery := "SELECT user_id, user_name, password FROM users WHERE user_name = $1"

    var user models.User
    err := db.QueryRow(sqlQuery, requestUser.UserName).Scan(&user.UUID, &user.UserName, &user.Password)
    if (err != nil) {
        log.Printf("Unable to query database %v", err)
        return user, err
    }

    return user, nil
}

// CheckUserNameQuery queries the database for the number of users with the provided user name.
// Used to check if a user name is taken or not.
func CheckUserNameQuery(userName string) bool {
    db := createConnection()
    defer db.Close()

    var nrOfRows int64
    sqlQuery := "SELECT COUNT(*) FROM users WHERE user_name = $1"
    err := db.QueryRow(sqlQuery, userName).Scan(&nrOfRows)
    if (err != nil) {
        log.Fatalf("Unable to query database %v", err)
    }

    if (nrOfRows >= 1) {
        return true
    } 

    return false
}

//InsertUserQuery inserts a user into the database with the given user model.
func InsertUserQuery(user models.User) uuid.UUID {
    db := createConnection()
    defer db.Close()

    sqlQuery := "INSERT INTO users (user_id, user_name, password) VALUES ($1, $2, $3) RETURNING user_id"

    var userId uuid.UUID

    err := db.QueryRow(sqlQuery, user.UUID, user.UserName, user.Password).Scan(&userId)
    if (err != nil) {
        log.Fatalf("Unable to insert user %v", err)
    }

    return userId
}

