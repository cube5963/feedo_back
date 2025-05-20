package create

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"feedo_back/middleware/db"
	"feedo_back/middleware/db/formdb"
	"feedo_back/middleware/models"
)

type FormResponse struct {
	Message string `json:"message" example:"Form created successfully"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Failed to create form"`
}

// CreateForm endpoint
// @Summary Create Form
// @Description Create a new form
// @Tags form
// @Accept json
// @Produce json
// @Param form body models.CreateFormRequest true "Form data"
// @Success 200 {object} FormResponse "Form creation success"
// @Failure 500 {object} ErrorResponse "Form creation failure"
// @Router /form/create [post]
func CreateHandler(c *gin.Context) {
	var credentials models.CreateFormRequest

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid request payload"})
		return
	}

	database, err := db.GetDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to connect to the database"})
		return
	}

	success := formdb.CreateForm(database, &credentials)

	if success {
		c.JSON(http.StatusOK, FormResponse{"Form created successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to create form"})
	}

}
