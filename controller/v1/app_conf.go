package v1

import (
	"gtank/middleware/resp"
	"gtank/valid"

	"github.com/gin-gonic/gin"
)

type AppConf struct{}

func (AppConf) EnvAdd(c *gin.Context) {
	p := valid.EnvAddParam{}
	err := valid.BindJsonAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, p)
}

func (AppConf) EnvList(c *gin.Context) {

}

func (AppConf) EnvDel(c *gin.Context) {

}
