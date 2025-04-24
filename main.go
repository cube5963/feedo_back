package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "feedo_back/docs"

	"feedo_back/middleware/db"
)

// @title Feedo API
// @version 1.0
// @description Feedo API documentation
// @host localhost:5000
// @BasePath /

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()

	RegisterRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//dsn := "host=localhost user=postgres password=pass dbname=ojousama port=5432 sslmode=disable"
	//dsn := os.Getenv("DB_DSN")

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_NAME")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",host, user, password, dbName, port)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.Migrate(database)

	if err != nil {
		panic("failed to connect to database")
	}

	r.Run(":5000")
}
