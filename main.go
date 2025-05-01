package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

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

	database, err := db.GetDB()
    if err != nil {
        panic(fmt.Sprintf("failed to connect to database: %v", err))
    }

	db.Migrate(database)

	r.Run(":5000")
}