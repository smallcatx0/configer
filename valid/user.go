package valid

import (
	"gtank/middleware/resp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// token中携带的数据
type JWTData struct {
	Uid      int    `json:"uid,omitempty"`
	User     string `json:"user,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Truename string `json:"rname,omitempty"`
}

type Claim struct {
	*jwt.RegisteredClaims
	JWTData
}

func (j *JWTData) Generate() (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)
	c := &Claim{
		RegisteredClaims: &jwt.RegisteredClaims{},
	}
	// 设置过期时间 3600s
	c.ExpiresAt = jwt.NewNumericDate(time.Now().Add(5 * time.Minute))
	c.JWTData = *j
	t.Claims = c
	return t.SignedString([]byte("sk"))
}

func JWTPase(token string) (*Claim, error) {
	t, err := jwt.ParseWithClaims(token, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("sk"), nil
	})
	if err != nil {
		if e, ok := err.(jwt.ValidationError); ok {
			if e.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, resp.LoginTimeOut
			} else {
				return nil, resp.IllegalToken
			}
		}
		return nil, err
	}
	data, ok := t.Claims.(*Claim)
	if ok && t.Valid {
		return data, nil
	}
	return nil, resp.IllegalToken
}

// 获取jwt中的用户信息
func UserInfo(c *gin.Context) (*JWTData, bool) {
	data, ok := c.Get("jwtinfo")
	if !ok {
		// 解析
		return UserInfoPase(c)
	}
	ret, ok := data.(JWTData)
	return &ret, ok
}

func UserInfoPase(c *gin.Context) (*JWTData, bool) {
	tokenStr := strings.TrimSpace(c.GetHeader("Authorization"))
	if tokenStr == "" {
		return nil, false
	}
	raw, err := JWTPase(tokenStr)
	if err != nil {
		return nil, false
	}
	c.Set("jwtinfo", raw.JWTData)
	return &raw.JWTData, true
}
