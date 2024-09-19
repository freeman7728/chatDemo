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
	if o.From != id {
		resp.Msg = "不可操作其他用户"
		resp.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	relationServ := service.GetRelationServIns()
	resp = relationServ.CreateRelation(o)
	ctx.JSON(http.StatusOK, resp)
	return
}
