package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
	_ "feedo_back/docs"
)

// @title Feedo API
// @version 1.0
// @description Feedo API documentation
// @host localhost:5000
// @BasePath /

func main() {
	r := gin.Default()

	RegisterRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":5000")
}