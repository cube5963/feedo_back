package register

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"feedo_back/middleware/db"
	"feedo_back/middleware/models"
	"feedo_back/middleware/db/userdb"
)

// User register endpoint
// @Summary User Registration
// @Description Register a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.UserRequest true "User data"
// @Success 200 {object} models.SuccessResponse "ユーザー登録成功"
// @Failure 400 {object} models.ErrorResponse "リクエスト不正"
// @Failure 500 {object} models.ErrorResponse "ユーザー登録失敗"
// @Router /user/register [post]
func RegisterHandler(c *gin.Context) {
	/*var credentials struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}*/

	var credentials models.UserRequest

	if err := c.ShouldBindJSON(&credentials); err != nil {
		//c.JSON(http.StatusBadRequest, models.ErrorResponse{"Invalid request payload"})
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid request payload"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error:"failed to hash password"})
        return
    }

	credentials.Password = string(hashedPassword)
	
	database, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to connect to the database"})
		return
	}

	// Attempt to create the user in the database
	success := userdb.RegisterUser(database, &credentials)

	// Record the result of the operation
	if success {
		c.JSON(http.StatusOK, models.SuccessResponse{Message: "User registered successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to register user"})
	}
}