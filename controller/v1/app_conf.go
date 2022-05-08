package v1

import (
	"gtank/middleware/resp"
	"gtank/models/dao/modb"
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
	// 准备数据
	d := modb.AppEnv{
		Name: p.Name,
		Type: p.Type,
		Sign: p.Sign,
		Desc: p.Desc,
		Principal: modb.Contact{
			Name:  p.OwnerName,
			Phone: p.OwnerPhone,
		},
	}
	// 先查再写

	id, err := d.Save()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, id)
}

func (AppConf) EnvList(c *gin.Context) {

}

func (AppConf) EnvDel(c *gin.Context) {

}
