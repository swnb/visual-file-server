package controllers

import (
	"github.com/gin-gonic/gin"
)

type ginHandler = func(c *gin.Context)

// Controller define controller properties
type Controller struct {
	Group       interface{} // group type string or []string
	URLPath     string
	Method      interface{} // methods type string or []string
	Handler     ginHandler  // center handler
	Middlewares interface{} // middlewares type []ginHandler or ginHandler
	Logger      interface{} // 预留的字段
	Config      interface{} // 预留的字段
}

// CreateControllerFunc create controller
type CreateControllerFunc = func() Controller

// init create handler fn with cap 100
var handlers = make([]CreateControllerFunc, 0, 100)

// GetControllers get all controllers
func GetControllers() []CreateControllerFunc {
	return handlers
}

func addController(handler ...CreateControllerFunc) {
	handlers = append(handlers, handler...)
}

// GroupController 组路由的控制器
type GroupController struct {
	Group   string
	Handler func(*gin.Context)
}

// CreateGroupControllerFunc create group controller
type CreateGroupControllerFunc = func() GroupController

// init group handlers
var groupHandlers = []CreateGroupControllerFunc{}

// GetGroupControllers get all group controllers
func GetGroupControllers() []CreateGroupControllerFunc {
	return groupHandlers
}

// append group controller
func addGroupController(fn ...CreateGroupControllerFunc) {
	groupHandlers = append(groupHandlers, fn...)
}

// alias map[string] interface j
type j = map[string]interface{}
