package sgcjwt

import (
	"errors"

	sgHttp "github.com/Wolf-Jia/salmon-go/http"

	"github.com/gin-gonic/gin"
)

var (
	//KeyTokenClaim token写入gin.context的名称
	KeyTokenClaim = "token-claims"
)

//Middleware token校验中间件
func Middleware(tokenName string, secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if tokenName == "" {
			tokenName = "Authorization"
		}
		token := c.Request.Header.Get(tokenName)
		if token == "" {
			sgHttp.Error(c, 401, errors.New("未认证用户"), "token无效")
			c.Abort()
			return
		}
		claims, err := DecrypToken(token, secretKey)
		if err != nil {
			sgHttp.Error(c, 401, err, "token校验失败")
			c.Abort()
			return
		}
		c.Set(KeyTokenClaim, claims)
	}
}
