package middleware

import (
	"chat/pkg/utils"
	"chat/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ParseToken(c *gin.Context) {
	//if !conf.EnableJwt {
	//	c.Next()
	//	return
	//}
	_, err := utils.ParseToken(c.GetHeader("Authorization"))
	fmt.Println(c.GetHeader("Authorization"), 1111)
	if err != nil {
		c.JSON(http.StatusForbidden, serializer.Response{
			Status: http.StatusForbidden,
			Msg:    "登录状态过期，请重新登录",
		})
		c.Abort()
		return
	}
	c.Next()
}
