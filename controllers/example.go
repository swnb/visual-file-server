package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	addController(example)
}

func example() Controller {

	handler := func(c *gin.Context) {
		responseBody := j{"code": 0, "message": "success"}
		c.Set("response body", responseBody)
		c.JSON(http.StatusOK, responseBody)
	}

	return Controller{
		Group:   []string{"/v1"},
		URLPath: "/alive/",
		Method:  "GET",
		Handler: handler,
	}
}
