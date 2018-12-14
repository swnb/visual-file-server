package router

import (
	"log"
	"strings"
	"visual-file-server/controllers"

	"github.com/gin-gonic/gin"
)

func registerMethods(groupRouter *gin.RouterGroup, method, URLPath string, handler func(*gin.Context)) {
	switch strings.ToLower(method) {
	case "get":
		groupRouter.GET(URLPath, handler)
	case "post":
		groupRouter.POST(URLPath, handler)
	case "delete":
		groupRouter.DELETE(URLPath, handler)
	case "put":
		groupRouter.PUT(URLPath, handler)
	case "options":
		groupRouter.OPTIONS(URLPath, handler)
	case "head":
		groupRouter.HEAD(URLPath, handler)
	default:
		log.Panic("error in router method ", method)
	}
}

func setHandler(groupPath, URLPath string, methods interface{}, handler func(*gin.Context)) {
	var groupRouter *gin.RouterGroup
	if v, ok := groupRouterMap[groupPath]; ok {
		groupRouter = v
	} else if groupPath == "" {
		groupRouter = groupRouterMap["/"]
	} else {
		groupRouter = registGroupRouter(groupPath, nil)
	}

	switch value := methods.(type) {
	case string:
		registerMethods(groupRouter, value, URLPath, handler)
	case []string:
		for _, v := range value {
			registerMethods(groupRouter, v, URLPath, handler)
		}
	default:
		log.Panic("error in regist method type ", methods)
	}
}

func registHandler() {
	// 注册开始控制器函数
	for _, constructor := range controllers.GetControllers() {
		controller := constructor()

		groupPaths := controller.Group
		URLPath := controller.URLPath
		methods := controller.Method
		handler := controller.Handler

		// 解析 groupPath 的数据
		switch value := groupPaths.(type) {
		case string:
			setHandler(value, URLPath, methods, handler)
		case []string:
			for _, v := range value {
				setHandler(v, URLPath, methods, handler)
			}
		case nil:
			setHandler("/", URLPath, methods, handler)
		default:
			log.Panic("type of Group must be []string or string")
		}
	}
}
