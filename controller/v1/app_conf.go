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
		"id": id,
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
		"id": id,
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

func (AppConf) FileAdd(c *gin.Context) {
	p := valid.AppFileAddParam{}
	err := valid.BindJsonAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	// 查询app信息，env信息
	e := &modb.AppEnv{}
	err = e.GetbySign(p.EnvSign)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	a := &modb.App{}
	err = a.GetbySign(p.AppSign)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	file := modb.AppConf{
		Env: modb.EnvBaseInfo{
			Type:    e.Type,
			Sign:    e.Sign,
			Name:    e.Name,
			Contact: e.Principal.Name,
		},
		App: modb.AppBaseInfo{
			Name:    a.Name,
			Sign:    a.Sign,
			Contact: a.Principal.Name,
		},
		Principal: modb.Contact{
			Name:  p.Header,
			Phone: p.HeaderPhone,
		},
		File: modb.ConfFile{
			Name:    p.FileName,
			Content: p.Content,
			Type:    p.Type,
		},
	}
	id, err := file.SaveOne()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	// 写入app 文件信息
	appfiles := modb.AppFile{
		AppSign: p.AppSign,
		EnvSign: p.EnvSign,
	}
	err = appfiles.AddFile(p.FileName)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, map[string]interface{}{
		"id": id,
	})
}

func (AppConf) History(c *gin.Context) {
	p := valid.AppFileParam{}
	err := valid.BindQueryAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	f := modb.AppConf{}
	list, err := f.FileHistory(p.Env, p.App, p.File)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, list)
}

func (AppConf) Top(c *gin.Context) {
	p := valid.AppConfParam{}
	err := valid.BindQueryAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	app := &modb.AppFile{
		AppSign: p.App,
		EnvSign: p.Env,
	}
	err = app.GetFiles()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	confs := make([]modb.AppConf, 0, len(app.Files))
	confapp := modb.AppConf{}
	// 挨个查询配置文件
	for n, s := range app.Files {
		if s != modb.FileStatus_Enable {
			continue
		}
		conf, err := confapp.Top(app.EnvSign, app.AppSign, n)
		if err != nil {
			resp.Fail(c, err)
			return
		}
		confs = append(confs, *conf)
	}
	resp.Succ(c, confs)
}

func (AppConf) FileDel(c *gin.Context) {
	p := valid.AppFileParam{}
	err := valid.BindJsonAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	app := modb.AppFile{
		EnvSign: p.Env,
		AppSign: p.App,
	}
	err = app.DelFile(p.File)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, nil)
}
