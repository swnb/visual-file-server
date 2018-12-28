package main

import (
	"visual-file-server/middleware"
	"visual-file-server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(
		middleware.Cors,
		middleware.Logger,
	)

	r = router.InitRouter(r)

	r.Run(":80")
}
