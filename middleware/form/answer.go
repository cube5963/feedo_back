package form

import (
	"feedo_back/middleware/db"
	"feedo_back/middleware/db/formdb"
	"feedo_back/middleware/models"
	"feedo_back/middleware/user/login"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAnswerHandler(c *gin.Context){
	var credentials models.CreateAnswerRequest

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	isValid := login.Verify(credentials.Jwt)

	if !isValid {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid JWT"})
		return
	}

	database, err := db.GetDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to connect to the database"})
		return
	}

	success := formdb.CreateAnswer(database, &credentials)

	if success {
		c.JSON(http.StatusOK, models.SuccessResponse{Message: "Answer created successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create answer"})
	}
}