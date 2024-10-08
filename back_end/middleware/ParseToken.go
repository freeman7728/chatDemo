package middleware

import (
	"chat/pkg/utils"
	"chat/serializer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ParseTokenForWebsocket(c *gin.Context) {
	//if !conf.EnableJwt {
	//	c.Next()
	//	return
	//}
	claim, err := utils.ParseToken(c.Query("token"))
	if err != nil {
		c.JSON(http.StatusForbidden, serializer.Response{
			Status: http.StatusForbidden,
			Msg:    "登录状态过期，请重新登录",
		})
		c.Abort()
		return
	}
	c.Set("id", claim.Id)
	c.Next()
}

func ParseToken(c *gin.Context) {
	//if !conf.EnableJwt {
	//	c.Next()
	//	return
	//}
	claim, err := utils.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusForbidden, serializer.Response{
			Status: http.StatusForbidden,
			Msg:    "登录状态过期，请重新登录",
		})
		c.Abort()
		return
	}
	c.Set("id", claim.Id)
	c.Next()
}
