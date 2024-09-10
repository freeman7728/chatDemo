package middleware

import (
	"chat/conf"
	"chat/pkg/utils"
	"chat/serializer"
	"github.com/gin-gonic/gin"
)

func ParseToken(c *gin.Context) {
	if !conf.EnableJwt {
		c.Next()
		return
	}
	_, err := utils.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(500, serializer.Response{
			Status: 500,
			Msg:    "登录状态过期，请重新登录",
		})
		c.Abort()
		return
	}
	c.Next()
}
