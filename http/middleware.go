package sgchttp

import "github.com/gin-gonic/gin"

// CrossOrigin is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
// 解决跨域问题
func CrossOrigin(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Content-Type", "application/json")
	//放行所有OPTIONS方法
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
	}
	c.Next()
}
