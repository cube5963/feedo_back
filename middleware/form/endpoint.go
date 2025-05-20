package form

import (
	"github.com/gin-gonic/gin"
)

func RootForm(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Form endpoint reached!",
	})
}

func Form(r *gin.Engine) {
	formGroup := r.Group("/form")
	{
		formGroup.GET("/", RootForm)
	}
}