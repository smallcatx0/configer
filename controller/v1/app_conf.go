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
	id, err := d.NewOne()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, map[string]interface{}{
		"ID": id,
	})
}

func (AppConf) EnvEdit(c *gin.Context) {
	p := valid.EnvEditParam{}
	err := valid.BindJsonAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	d := modb.AppEnv{
		Sign: p.Sign,
		Name: p.Name,
		Desc: p.Desc,
		Principal: modb.Contact{
			Name:  p.OwnerName,
			Phone: p.OwnerPhone,
		},
	}
	count, err := d.Edit()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, map[string]interface{}{
		"count": count,
	})
}

func (AppConf) EnvList(c *gin.Context) {
	d := modb.AppEnv{}
	ret, err := d.List()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, ret)
}

func (AppConf) EnvDel(c *gin.Context) {
	p := struct {
		Sign string `json:"sign" binding:"required"`
	}{}
	err := valid.BindJsonAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	d := modb.AppEnv{
		Sign: p.Sign,
	}
	count, err := d.Del()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, map[string]interface{}{
		"count": count,
	})
}

func (AppConf) AppAdd(c *gin.Context) {
	p := valid.AppAddParam{}
	err := valid.BindJsonAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	a := modb.App{
		Name: p.Name,
		Sign: p.Sign,
		Desc: p.Desc,
		Principal: modb.Contact{
			Name:  p.OwnerName,
			Phone: p.OwnerPhone,
		},
	}
	id, err := a.NewOne()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, map[string]interface{}{
		"ID": id,
	})
}

func (AppConf) AppEdit(c *gin.Context) {
	p := valid.AppEditParam{}
	err := valid.BindJsonAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	a := modb.App{
		Sign: p.Sign,
		Name: p.Name,
		Principal: modb.Contact{
			Name:  p.OwnerName,
			Phone: p.OwnerPhone,
		},
	}
	count, err := a.Edit()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, map[string]interface{}{
		"count": count,
	})
}

func (AppConf) AppList(c *gin.Context) {
	a := modb.App{}
	ret, err := a.List()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, ret)
}

func (AppConf) AppDel(c *gin.Context) {
	p := struct {
		Sign string `json:"sign" binding:"required"`
	}{}
	err := valid.BindJsonAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	a := modb.App{
		Sign: p.Sign,
	}
	count, err := a.Del()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, map[string]interface{}{
		"count": count,
	})
}
