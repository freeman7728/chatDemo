package api

import (
	"chat/model"
	"chat/serializer"
	"chat/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGroupApi(ctx *gin.Context) {
	groupServ := service.GetGroupServIns()
	var m model.Group
	var resp serializer.Response
	resp.Status = http.StatusOK
	err := ctx.ShouldBindBodyWithJSON(&m)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp = groupServ.CreateGroupServ(m)
	ctx.JSON(http.StatusOK, resp)
	return
}
