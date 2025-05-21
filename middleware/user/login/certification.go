package login

import (
	"feedo_back/middleware/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


// User verify endpoint
// @Summary User verify endpoint
// @Description This is a user verify endpoint
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.UserVerifyRequest true "User verify data"
// @Success 200 {object} models.VerifyResponse "Returns verification result"
// @Router /user/verify [post]
func VerifyHandler(c *gin.Context) {
	/*var credentials struct {
		Token string `json:"token"`
	}*/

	var credentials models.UserVerifyRequest
	
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	c.JSON(http.StatusOK, models.VerifyResponse{Verify: Verify(credentials.Token)})
}