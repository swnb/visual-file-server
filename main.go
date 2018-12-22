package main

import (
	"visual-file-server/middleware"
	"visual-file-server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r = router.InitRouter(r)

	r.Use(middleware.Logger)

	r.Run(":80")
}
