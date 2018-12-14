package router

import (
	"github.com/gin-gonic/gin"
)

var ginRouter *gin.Engine

// InitRouter 初始化所有的路由
func InitRouter(r *gin.Engine) *gin.Engine {

	// 初始化的gin引擎
	ginRouter = r

	// 初始化组路由
	initGroupRouters()

	// 注册控制器
	registHandler()

	return ginRouter
}
