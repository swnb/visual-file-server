package main

import (
	"visual-file-server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r = router.InitRouter(r)

	r.Run(":80")
}
