package main

import (
	"log"
	"os"

	"cooding/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	router := gin.Default()
	router.GET("/me", controllers.GetInfo)
	router.Run(":" + port)
}
