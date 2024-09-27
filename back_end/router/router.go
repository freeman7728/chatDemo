package router

import (
	"chat/api"
	"chat/middleware"
	"chat/model"
	"chat/serializer"
	"chat/service"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	v1 := r.Group("/")
	{
		v1.POST("/register", api.UserApi{}.UserRegister)
		v1.POST("/login", api.UserApi{}.UserLogin)
		//获取升级连接的请求
		//v1.GET("/ws", middleware.ParseToken, middleware.CheckUserMiddleware, service.Handler)
		v1.GET("/ws", middleware.ParseTokenForWebsocket, middleware.CheckUserMiddleware, service.Handler)
	}
	v2 := r.Group("/student")
	{
		CRUD[model.Student](v2, ":id")
	}
	relation := r.Group("/relation")
	{
		relation.POST("/create", middleware.ParseToken, api.CreateRelationHandler)
		relation.GET("/get-relation-list", middleware.ParseToken, api.GetRelationHandler)
		relation.POST("del", middleware.ParseToken, api.DelRelationHandler)
	}
	testJwt := r.Group("/")
	{
		testJwt.GET("/ping", middleware.ParseToken, func(c *gin.Context) {
			id, _ := c.Get("id")
			var resp serializer.Response
			resp.Data = id
			c.JSON(200, resp)
		})
	}
	return r
}

func CRUD[T any](group *gin.RouterGroup, idParam string) *gin.RouterGroup {
	group.POST("", api.InsertHandler[T])
	return group
}
