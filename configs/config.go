package configs

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}
