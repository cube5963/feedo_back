package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"feedo_back/middleware/db"
	"feedo_back/middleware/db/userdb"
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
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} LoginResponse "Returns JWT token"
// @Router /user/login [post]
func LoginHandler(c *gin.Context) {
	// Parse login credentials from request body
	var credentials struct {
		Email string `json:"email"`
		Password string `json:"password"`
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

	success, userPassword := userdb.LoginUser(database, credentials.Email)
	// Authenticate user (this is a placeholder, replace with real authentication logic)
	if !success {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
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
