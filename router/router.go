package router

import (
	"chat/api"
	"chat/middleware"
	"chat/service"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/")
	{
		v1.POST("/register", api.UserApi{}.UserRegister)
		v1.GET("/login", api.UserApi{}.UserLogin)
		//获取升级连接的请求
		v1.GET("/ws", middleware.ParseToken, middleware.CheckUserMiddleware, service.Handler)
	}
	testJwt := r.Group("/")
	{
		testJwt.GET("/ping", middleware.ParseToken, func(c *gin.Context) {
			c.JSON(200, "success")
		})
	}
	return r
}
