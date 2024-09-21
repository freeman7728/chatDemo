package api

import (
	"chat/model"
	"chat/serializer"
	"chat/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type relationApi struct {
}

func CreateRelationHandler(ctx *gin.Context) {
	var resp serializer.Response
	var o model.Relation
	if err := ctx.ShouldBindBodyWithJSON(&o); err != nil {
		resp.Error = "参数绑定错误"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	id, _ := ctx.Get("id")
	//类型断言
	intValue, _ := id.(int64)
	o.Source = intValue
	relationServ := service.GetRelationServIns()
	resp = relationServ.CreateRelation(o)
	ctx.JSON(http.StatusOK, resp)
	return
}

func DelRelationHandler(ctx *gin.Context) {
	var resp serializer.Response
	var o model.Relation
	if err := ctx.ShouldBindBodyWithJSON(&o); err != nil {
		resp.Error = "参数绑定错误"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	id, _ := ctx.Get("id")
	//类型断言
	intValue, _ := id.(int64)
	o.Source = intValue
	relationServ := service.GetRelationServIns()
	resp = relationServ.DelRelation(o)
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetRelationHandler(ctx *gin.Context) {
	var resp serializer.Response
	temp, _ := ctx.Get("id")
	id := temp.(int64)
	relationServ := service.GetRelationServIns()
	resp = relationServ.GetRelation(id)
	ctx.JSON(http.StatusOK, resp)
}
