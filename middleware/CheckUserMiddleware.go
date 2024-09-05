package middleware

import (
	"chat/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CheckUserMiddleware TODO 创建用户验证中间件
func CheckUserMiddleware(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	var u model.User
	u.ID = uint(uid)
	if u.CheckUid() {
		c.Next()
		return
	}
	//TODO 收到了错误信息，但是在body里面，postman的websocket升级请求不显示返回的body
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid uid"})
	c.Abort()
}
