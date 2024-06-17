package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowEndOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "SHOW Opening",
	})
}
