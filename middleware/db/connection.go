package db

import (
    "fmt"
    "os"
    "sync"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var (
    DB *gorm.DB
    once sync.Once
)

func GetDB() (*gorm.DB, error) {
    var err error
    once.Do(func() {
        host := os.Getenv("POSTGRES_HOST")
        user := os.Getenv("POSTGRES_USER")
        password := os.Getenv("POSTGRES_PASSWORD")
        dbName := os.Getenv("POSTGRES_NAME")
        port := os.Getenv("POSTGRES_PORT")

        dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    })

    return DB, err
}