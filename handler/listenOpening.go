package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiagolua/gonow.git/schemas"
)

func ListenEndOpeninnHandler(ctx *gin.Context) {
	openings := []schemas.Product{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
