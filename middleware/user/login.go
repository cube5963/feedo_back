package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler handles user login and returns a JWT token
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
