package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rebound.sh/gin-setup/internal/middleware"
)

type SampleData struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.2"})

	// Sample middleware
	router.Use(middleware.Logger)

	// Retrieves (would be) needed data/environment
	router.Use(middleware.SetEnv())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, SampleData{Foo: "f", Bar: "b"})
	})

	router.Handler()	

	router.Run("localhost:8080")
}