package router

import (
	"visual-file-server/controllers"

	"github.com/gin-gonic/gin"
)

var groupRouterMap = make(map[string]*gin.RouterGroup)

func registGroupRouter(groupPath string, groupMiddleWare func(*gin.Context)) *gin.RouterGroup {

	groupRouter := ginRouter.Group(groupPath)

	groupRouterMap[groupPath] = groupRouter

	if groupMiddleWare != nil {
		groupRouter.Use(groupMiddleWare)
	}

	return groupRouter
}

func initGroupRouters() {

	// 默认的组路由
	registGroupRouter("/", nil)
	for _, constructor := range controllers.GetGroupControllers() {
		groupController := constructor()

		groupPath := groupController.Group
		groupMiddleWare := groupController.Handler

		// 注册组路由和控制器
		groupRouterMap[groupPath] = registGroupRouter(groupPath, groupMiddleWare)
	}
}
