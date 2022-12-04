package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetEnv() gin.HandlerFunc {
	log.Println("[SET ENVIRONMENT]")

	return func (c *gin.Context)  {
		c.Next()	
	}
}
