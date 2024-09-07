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
		//获取升级连接的请求
		v1.GET("/ws", middleware.ParseToken, middleware.CheckUserMiddleware, service.Handler)
	}
	v2 := r.Group("/")
	{
		v2.GET("/login", api.UserApi{}.UserLogin)
	}
	testJwt := r.Group("/")
	{
		testJwt.GET("/ping", middleware.ParseToken, func(c *gin.Context) {
			c.JSON(200, "success")
		})
	}
	return r
}
