package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerifyResponse struct {
	Verify bool `json:"verify" example:"true"`
}

// User verify endpoint
// @Summary User verify endpoint
// @Description This is a user verify endpoint
// @Tags user
// @Accept json
// @Produce json
// @Param token body string true "Token"
// @Success 200 {object} VerifyResponse "Returns verification result"
// @Router /user/verify [post]
func VerifyHandler(c *gin.Context) {
	var credentials struct {
		Token string `json:"token"`
	}
	
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"verify": verify(credentials.Token)})
}