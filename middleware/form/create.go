package form

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"feedo_back/middleware/db"
	"feedo_back/middleware/db/formdb"
	"feedo_back/middleware/models"
	"feedo_back/middleware/user/login"
)

// CreateForm endpoint
// @Summary Create Form
// @Description Create a new form
// @Tags form
// @Accept json
// @Produce json
// @Param form body models.CreateFormRequest true "Form data"
// @Success 200 {object} models.SuccessResponse "Form creation success"
// @Failure 500 {object} models.ErrorResponse "Form creation failure"
// @Router /form/create [post]
func CreateHandler(c *gin.Context) {
	var credentials models.CreateFormRequest

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

	success := formdb.CreateForm(database, &credentials)

	if success {
		c.JSON(http.StatusOK, models.SuccessResponse{Message: "Form created successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create form"})
	}

}
