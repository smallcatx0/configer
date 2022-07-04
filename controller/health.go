package controller

import (
	"gtank/bootstrap"
	"gtank/middleware/resp"

	"github.com/gin-gonic/gin"
)

type Health struct{}

func (Health) Healthz(c *gin.Context) {
	resp.Succ(c, "")
}

func (Health) Ready(c *gin.Context) {
	resp.Succ(c, "ooooook")
}

// 重新加载配置文件
func (Health) ReloadConf(c *gin.Context) {
	bootstrap.InitConf(&bootstrap.Param.C)
	bootstrap.InitLog()
	bootstrap.InitDB()
	resp.Succ(c, "yeeeeee")
}

func (Health) Test(c *gin.Context) {

}
