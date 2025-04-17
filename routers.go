package main

import (
    "github.com/gin-gonic/gin"

	"feedo_back/middleware/user"
)

func RegisterRoutes(r *gin.Engine) {
	user.User(r)
}