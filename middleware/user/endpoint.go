package user

import (
	"github.com/gin-gonic/gin"
	"feedo_back/middleware/user/login"
	"feedo_back/middleware/user/register"
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
		userGroup.POST("/login", login.LoginHandler)
		userGroup.POST("/verify", login.VerifyHandler)
		userGroup.POST("/register", register.RegisterHandler)
	}
}
