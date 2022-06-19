package v1

import (
	"gtank/middleware/resp"
	"gtank/models/dao/modb"
	"gtank/valid"

	"github.com/gin-gonic/gin"
)

type DbCron struct{}

func (DbCron) TTLAdd(c *gin.Context) {
	p := valid.DbTTLParam{}
	err := valid.BindJsonAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	v := modb.DbTTL{
		DSN:   p.DSN,
		DB:    p.DB,
		Table: p.Table,
		Field: p.Field,
		Cron:  p.Cron,
		TTL:   p.TTL,
		Limit: p.Limit,
	}
	ok, err := v.IsExist()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	if ok {
		resp.Fail(c, resp.Repeat)
		return
	}
	id, err := v.NewOne()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, map[string]interface{}{
		"id": id,
	})
}
