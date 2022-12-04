package main

import (
	"github.com/gin-gonic/gin"
	"rebound.sh/gin-setup/internal/controllers"
	"rebound.sh/gin-setup/internal/middleware"
)



func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.2"})

	// Sample middleware
	router.Use(middleware.Logger)

	// Retrieves (would be) needed data/environment
	router.Use(middleware.SetEnv())

	router.GET("/", controllers.RootController)

	router.Handler()	

	router.Run("localhost:8080")
}