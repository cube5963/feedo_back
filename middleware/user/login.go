package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginResponse struct {
    Token string `json:"token" example:"eyJhbGciOijiwzIJNviIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJpc3N1ZXIiLCJhsIiOiJzdWJqZWN0IiwiYXVkIjpbImF1ZGllbmNlIsdvssImV4cCI6MTMIJNTQ2NzM0MywibmJmIjoxNzQ1NhtbMjgzLCJpYXQiOjE3NDU0NjcyODMsImp0aSI6ImlkIn0.dxgM6uH2F8ZglV_xcPhjCRnOSJBYq9oeS1TDLkLg_eg"`
}

// User login endpoint
// @Summary User login endpoint
// @Description This is a user login endpoint
// @Tags user
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} LoginResponse "Returns JWT token"
// @Router /user/login [post]
func LoginHandler(c *gin.Context) {
	// Parse login credentials from request body
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Authenticate user (this is a placeholder, replace with real authentication logic)
	if credentials.Username != "admin" || credentials.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the token
	c.JSON(http.StatusOK, gin.H{"token": token})
}
