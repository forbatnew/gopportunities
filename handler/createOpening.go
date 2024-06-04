package handler

import (
	"net/http"

	"github.com/forbatnew/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /appi/v1

// CreateOpenings godoc
// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]

func CreateOpeningHandler(ctx *gin.Context) {

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating openings request: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating openings on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)

	// logger.Infof("request received: %+v", request)
}
