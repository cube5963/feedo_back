package form

import (
	"net/http"

	//"feedo_back/middleware/db"
	"feedo_back/middleware/models"
	"feedo_back/middleware/user/login"

	"github.com/gin-gonic/gin"
)

func EditFormHandler(c *gin.Context) {
	var credentials models.EditFormRootRequest

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	isValid := login.Verify(credentials.Jwt)

	if !isValid {
		c.JSON(http.StatusUnauthorized, ErrorResponse{"Invalid JWT"})
		return
	}

	//database, err := db.GetDB()
}