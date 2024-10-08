package api

import (
	"chat/model"
	"chat/serializer"
	"chat/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddGroupMemberApi(ctx *gin.Context) {
	var resp serializer.Response
	var m model.GroupMember
	err := ctx.ShouldBindBodyWithJSON(&m)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusOK, resp)
		return
	}
	g := service.GetGroupMemberServIns()
	resp = g.AddGroupMember(m)
	ctx.JSON(http.StatusOK, resp)
	return
}
