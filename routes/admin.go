package routes

import (
	v1 "gtank/controller/v1"

	"github.com/gin-gonic/gin"
)

// 内部管理系统接口
func registeAdmin(router *gin.RouterGroup) {
	routV1 := router.Group("/v1")
	appConf := &v1.AppConf{}
	{
		routV1.GET("/conf/envs", appConf.EnvList)
		routV1.POST("/conf/env-add", appConf.EnvAdd)
		routV1.POST("/conf/env-edit", appConf.EnvEdit)
		routV1.POST("/conf/env-del", appConf.EnvDel)

		routV1.GET("/conf/apps", appConf.AppList)
		routV1.POST("/conf/app-add", appConf.AppAdd)
		routV1.POST("/conf/app-edit", appConf.AppEdit)
		routV1.POST("/conf/app-del", appConf.AppDel)
	}

}
