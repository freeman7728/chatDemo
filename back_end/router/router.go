package router

import (
	"chat/api"
	"chat/middleware"
	"chat/model"
	"chat/serializer"
	"chat/service/WebsocketService"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	v1 := r.Group("/")
	{
		v1.POST("/register", api.UserApi{}.UserRegister)
		v1.POST("/login", api.UserApi{}.UserLogin)
		v1.GET("/ws", middleware.ParseTokenForWebsocket, middleware.CheckUserMiddleware, WebsocketService.Handler)
		v1.GET("/userWebsocket", middleware.ParseTokenForWebsocket, WebsocketService.UserWsHandler)
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
	authJwt := r.Group("/")
	{
		authJwt.GET("/ping", middleware.ParseToken, func(c *gin.Context) {
			id, _ := c.Get("id")
			var resp serializer.Response
			resp.Data = id
			c.JSON(200, resp)
		})
	}
	v3 := r.Group("/group")
	{
		v3.POST("/create", middleware.ParseToken, api.CreateGroupApi)
		v3.POST("/addMember", middleware.ParseToken, api.AddGroupMemberApi)
	}
	return r
}

func CRUD[T any](group *gin.RouterGroup, idParam string) *gin.RouterGroup {
	group.POST("", api.InsertHandler[T])
	return group
}
