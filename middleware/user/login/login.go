package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"feedo_back/middleware/db"
	"feedo_back/middleware/db/userdb"
	"feedo_back/middleware/models"
)

// User login endpoint
// @Summary User login endpoint
// @Description This is a user login endpoint
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.UserLoginRequest true "User login data"
// @Success 200 {object} models.LoginResponse "ログイン成功"
// @Failure 400 {object} models.ErrorResponse "リクエスト不正"
// @Failure 401 {object} models.ErrorResponse "認証失敗"
// @Failure 500 {object} models.ErrorResponse "サーバーエラー"
// @Router /user/login [post]
func LoginHandler(c *gin.Context) {
	// Parse login credentials from request body
	/*var credentials struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}*/

	var credentials models.UserLoginRequest

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	database, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to connect to the database"})
		return
	}

	success, userPassword := userdb.LoginUser(database, credentials.Email)
	// Authenticate user (this is a placeholder, replace with real authentication logic)
	if !success {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to generate token"})
		return
	}

	// Return the token
	c.JSON(http.StatusOK, models.LoginResponse{Token: token})
}
