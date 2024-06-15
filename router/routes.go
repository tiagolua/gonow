package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagolua/gonow.git/handler"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.CreatEndOpeningHandler)
		v1.POST("/opening", handler.ShowEndOpeningHandler)
		v1.DELETE("/opening", handler.DeleteEndOpeningHandler)
		v1.PUT("/opening", handler.ListenEndOpeninnHandler)
		v1.GET("/opening", handler.UpdateEndOpeningHandler)

	}
}
