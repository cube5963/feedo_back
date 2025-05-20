package create

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"feedo_back/middleware/db"
)

func CreateHandler(c *gin.Context) {
	var credentials struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	database, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
}