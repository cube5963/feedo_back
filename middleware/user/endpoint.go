package user

import (
	"github.com/gin-gonic/gin"
)


func RootUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User endpoint reached!",
	})
}

func User(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", RootUser)
		userGroup.POST("/login", LoginHandler)
	}
}
