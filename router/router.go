package router

import "github.com/gin-gonic/gin"

func Initializer() {
	//	initialize Router
	router := gin.Default()

	//Initialize Routes
	initializeRoutes(router)

	// Run the Server
	router.Run(":8080")
}
