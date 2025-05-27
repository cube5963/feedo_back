package form

import (
	"net/http"

	//"feedo_back/middleware/db"
	"feedo_back/middleware/user/login"
	"feedo_back/middleware/models"

	"github.com/gin-gonic/gin"
)

func EditFormHandler(c *gin.Context) {
	var credentials models.EditFormRootRequest

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	isValid := login.Verify(credentials.Jwt)

	if !isValid {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid JWT"})
		return
	}

	//database, err := db.GetDB()
}