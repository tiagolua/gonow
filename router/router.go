package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initializer() {
	//	initialize Router
	router := gin.Default()

	//Initialize Routes
	initializeRoutes(router)

	// Get the port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run the Server
	router.Run("0.0.0.0:" + port)
}
