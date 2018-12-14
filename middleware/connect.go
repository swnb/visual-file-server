package middleware

import (
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

var connectNum int64

// Count define the alive connect number
func Count(c *gin.Context) {
	atomic.AddInt64(&connectNum, 1)
	c.Next()
	atomic.AddInt64(&connectNum, -1)
}

// GetConnectNum atomic return current connect number
func GetConnectNum() int64 {
	return atomic.LoadInt64(&connectNum)
}
