package register

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"feedo_back/middleware/db"
	"feedo_back/middleware/models"
	"feedo_back/middleware/db/userdb"
)

type UserResponse struct {
    Message string `json:"message" example:"User registered successfully"`
}

type ErrorResponse struct {
    Error string `json:"error" example:"Failed to register user"`
}

// User register endpoint
// @Summary User Registration
// @Description Register a new user
// @Tags user
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} UserResponse "User registration success"
// @Failure 500 {object} ErrorResponse "User registration failure"
// @Router /user/register [post]
func RegisterHandler(c *gin.Context) {
	var credentials struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

	// Initialize user variable
	user := models.UserRequest{
		Name:     credentials.Name,
		Email:    credentials.Email,
		Password: string(hashedPassword),
	}
	
	database, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	// Attempt to create the user in the database
	success := userdb.RegisterUser(database, &user)

	// Record the result of the operation
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
	}
}