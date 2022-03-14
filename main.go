package main

import (
	"example.com/go-aldous-backend/controllers"

	"github.com/gin-gonic/gin"
)

func addControllerRoutes(router *gin.Engine) {
	controllers.AddAldousControllerRoutes(router)
}

func main() {
	router := gin.Default()
	addControllerRoutes(router)
	router.Run("localhost:8080")
}
