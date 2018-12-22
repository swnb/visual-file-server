package rules

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type j = gin.H

const (
	// CodeOK means everythings is ok
	CodeOK = 0
	// CodeError means reponse is not ok
	CodeError = 10000
	// CodeErrorQuery means query is error
	CodeErrorQuery = 10003
)

// Success response success operate
func Success(c *gin.Context, data j) (body j) {
	body = j{"code": 0, "message": "success", "data": data}
	c.JSON(http.StatusOK, body)
	return
}

// ErrorQuery response err that query of request is error
func ErrorQuery(c *gin.Context) (body j) {
	body = j{"code": CodeErrorQuery, "message": "error params"}
	c.JSON(http.StatusNotFound, body)
	return
}

// SelfDefineRes define response self
func SelfDefineRes(c *gin.Context, code int, message string) (body j) {
	body = j{"code": code, "message": message}
	c.JSON(http.StatusOK, body)
	return
}

// Nothing reponse nothing whit http code 204
func Nothing(c *gin.Context) (body j) {
	c.JSON(http.StatusNoContent, nil)
	return
}

// Error means some error happen,and we don't know how to do it
func Error(c *gin.Context) (body j) {
	body = j{
		"code":    "10099",
		"message": "unknow error",
	}
	c.JSON(http.StatusNotFound, body)
	return
}
