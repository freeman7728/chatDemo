package router

import (
	api2 "chat/api"
	"chat/middleware"
	"chat/model"
	"chat/service"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/")
	{
		v1.POST("/register", api2.UserApi{}.UserRegister)
		v1.GET("/login", api2.UserApi{}.UserLogin)
		//获取升级连接的请求
		v1.GET("/ws", middleware.ParseToken, middleware.CheckUserMiddleware, service.Handler)
	}
	v2 := r.Group("/student")
	{
		CRUD[model.Student](v2, ":id")
	}
	testJwt := r.Group("/")
	{
		testJwt.GET("/ping", middleware.ParseToken, func(c *gin.Context) {
			c.JSON(200, "success")
		})
	}
	return r
}

func CRUD[T any](group *gin.RouterGroup, idParam string) *gin.RouterGroup {
	group.POST("", api2.InsertHandler[T])
	return group
}
