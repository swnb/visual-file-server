package main

import (
	"golang-gin-template/middleware"
	"golang-gin-template/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.Logger)

	r = router.InitRouter(r)

	r.Run(":80")
}
