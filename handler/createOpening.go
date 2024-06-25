package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatEndOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.error())
		sendErrorf(ctx, http.StatusBadRequest, err.error())
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
		logger.Errorf("error creating opening: %v", err.error())
		sendError(ctx, http.StatusInternalServerError, "error creating on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
