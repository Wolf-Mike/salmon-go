package sgchttp

import (
	sgUtility "go-webpage-builder/salmon-go/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ok 通常成功数据处理
func Ok(c *gin.Context, data interface{}, msg string) {
	var res ResponseVto
	res.Data = data
	res.Msg = "成功" + msg
	res.RequestID = sgUtility.GenerateMsgIDFromContext(c)
	c.AbortWithStatusJSON(http.StatusOK, res.ReturnOk())
}

//Error204 简化版204错误处理
func Error204(c *gin.Context, err error, msg string) {
	Error(c, http.StatusNoContent, err, msg)
}

//Error400 简化版400错误处理
func Error400(c *gin.Context, err error, msg string) {
	if msg == "" {
		msg = "参数错误"
	}
	Error(c, http.StatusBadRequest, err, msg)
}

//Error 失败数据处理
func Error(c *gin.Context, code int, err error, msg string) {
	var res ResponseVto

	res.Msg = msg
	if err != nil {
		res.Msg = msg + ":" + err.Error()
	}

	res.RequestID = sgUtility.GenerateMsgIDFromContext(c)
	c.AbortWithStatusJSON(http.StatusOK, res.ReturnError(code))
}
