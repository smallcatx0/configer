package v1

import (
	"gtank/middleware/resp"
	"gtank/models/dao/modb"
	"gtank/valid"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DbCron struct{}

func (DbCron) TTLAdd(c *gin.Context) {
	p := valid.DbTTLAddParam{}
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
		Desc:  p.Desc,
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

func (DbCron) TTLEdit(c *gin.Context) {
	p := valid.DbTTLEditParam{}
	err := valid.BindJsonAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	id, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		resp.Fail(c, resp.ParamInValid("id 参数错误"))
		return
	}
	v := modb.DbTTL{ID: id}
	up := bson.D{}
	if p.Field != "" {
		up = append(up, bson.E{Key: "field", Value: p.Field})
	}
	if p.Cron != "" {
		up = append(up, bson.E{Key: "cron", Value: p.Cron})
	}
	if p.TTL != 0 {
		up = append(up, bson.E{Key: "ttl", Value: p.TTL})
	}
	if p.Limit != 0 {
		up = append(up, bson.E{Key: "limit", Value: p.Limit})
	}
	if p.Status != "" {
		up = append(up, bson.E{Key: "status", Value: p.Status})
	}
	if p.Desc != "" {
		up = append(up, bson.E{Key: "desc", Value: p.Desc})
	}
	count, err := v.Edit(up)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, map[string]interface{}{
		"count": count,
	})
}

func (DbCron) TTLDel(c *gin.Context) {
	p := struct {
		ID string `json:"id"`
	}{}
	err := valid.BindJsonAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	id, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		resp.Fail(c, resp.ParamInValid("id 参数错误"))
		return
	}
	v := modb.DbTTL{ID: id}
	err = v.Del()
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, nil)
}

func (DbCron) TTLList(c *gin.Context) {
	p := valid.DbTTLQParam{}
	err := valid.BindQueryAndCheck(c, &p)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	q := bson.D{}
	if p.DB != "" {
		q = append(q, primitive.E{Key: "db", Value: p.DB})
	}
	if p.Table != "" {
		q = append(q, primitive.E{Key: "table", Value: p.Table})
	}
	v := modb.DbTTL{}
	list, err := v.List(q, 0, 99999)
	if err != nil {
		resp.Fail(c, err)
		return
	}
	resp.Succ(c, list)
}
