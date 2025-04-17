package user

import (
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "User endpoint reached!",
			})
		})
		// Add login endpoint
		userGroup.POST("/login", LoginHandler)
	}
}
