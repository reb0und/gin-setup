package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// type middleware interface {
// 	Log()	string
// 	Error() error
// }

// type Logger struct {
// 	S string
// }

func Logger(c *gin.Context) {
	log.Println(c.Request.Method, "	| ", c.ClientIP())

	c.Next()
}