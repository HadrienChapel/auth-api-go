package utils

import (
    "auth-api/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&models.User{})
    DB = database
}
