package main

import (
    "log"
    "net/http"
    "os"

    "github.com/wille1101/plant-waterer/backend/router"

    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load(".env")
    if (err != nil) {
        log.Fatalf("Error loading .env file")
    }
    
    r := router.Router()
    log.Fatal(http.ListenAndServe(":" + os.Getenv("BACKEND_PORT"), r))
}
