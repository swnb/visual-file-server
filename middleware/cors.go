package middleware

import (
	"github.com/gin-gonic/gin"
)

// Cors set Header Access-Control-Allow-Origin = *
func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
