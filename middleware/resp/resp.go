package resp

import (
	"gtank/internal/conf"
	"gtank/pkg/glog"

	"github.com/gin-gonic/gin"
)

// RequestIDKey 唯一请求id
const RequestIDKey = "x-b3-traceid"

type body struct {
	ErrCode   int         `json:"errcode"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	RequestID string      `json:"request_id"`
}

func SuccJsonRaw(c *gin.Context, data string) {
	format := `{"errcode":%d,"msg":"%s","data":%s,"request_id":"%s"}`
	requestId := c.GetString(RequestIDKey)
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.String(200, format, 200, "操作成功", data, requestId)
}

func Response(c *gin.Context, err error) {
	var httpCode int
	b := body{}
	switch e := err.(type) {
	case *Exception:
		httpCode = e.HTTPCode
		b.ErrCode = e.ErrCode
		b.Msg = e.Msg
		b.Data = e.Data
	case Exception:
		httpCode = e.HTTPCode
		b.ErrCode = e.ErrCode
		b.Msg = e.Msg
		b.Data = e.Data
	default:
		httpCode = 500
		b.ErrCode = 50000
		if conf.Env() == "dev" {
			b.Msg = e.Error()
		} else {
			glog.Error(c.Request.RequestURI, e.Error())
			b.Msg = "服务错误"
		}
	}
	b.RequestID = c.GetHeader(RequestIDKey)
	c.JSON(httpCode, &b)
}

func Succ(c *gin.Context, data interface{}) {
	Response(c, NewSucc(data))
}

func Fail(c *gin.Context, err error) {
	Response(c, err)
}
