package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Logger log reponse
func Logger(c *gin.Context) {
	c.Next()
	if data, ok := c.Get("response"); ok {
		go fmt.Println(data)
	}
}
