package routes

import "github.com/gin-gonic/gin"

// 内部管理系统接口
func registeAdmin(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	appConf := v1.Group("/app-conf")
	{

	}
}
