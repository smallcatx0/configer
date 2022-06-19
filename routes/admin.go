package routes

import (
	v1 "gtank/controller/v1"

	"github.com/gin-gonic/gin"
)

// 内部管理系统接口
func registeAdmin(router *gin.RouterGroup) {
	confR := router.Group("/v1/conf")
	{
		confCtl := v1.AppConf{}
		confR.GET("/envs", confCtl.EnvList)
		confR.POST("/env-add", confCtl.EnvAdd)
		confR.POST("/env-edit", confCtl.EnvEdit)
		confR.POST("/env-del", confCtl.EnvDel)
		confR.GET("/apps", confCtl.AppList)
		confR.POST("/app-add", confCtl.AppAdd)
		confR.POST("/app-edit", confCtl.AppEdit)
		confR.POST("/app-del", confCtl.AppDel)
		confR.POST("/file-add", confCtl.FileAdd)
		confR.POST("/file-del", confCtl.FileDel)
		confR.GET("/file-history", confCtl.History)
		confR.GET("/appconf", confCtl.Top)
	}

	cronR := router.Group("/v1/dbcron")
	{
		dbcronCtl := v1.DbCron{}
		cronR.POST("/ttl-add", dbcronCtl.TTLAdd)
	}
}
