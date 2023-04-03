package database

import (
    "database/sql"
    _ "github.com/lib/pq"

    "log"
    "os"

    "github.com/joho/godotenv"
)

// createConnection creates the database connection. Used by all query functions.
func createConnection() *sql.DB {
    err := godotenv.Load(".env")
    if (err != nil) {
        log.Fatalf("Error loading .env file")

    }

    db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
    if (err != nil) {
        panic(err)
    }


    err = db.Ping()
    if (err != nil) {
        log.Fatalf("Unable to connect to database")
        panic(err)
    }

    return db
}
