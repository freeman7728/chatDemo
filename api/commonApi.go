package api

import (
	"chat/serializer"
	"chat/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertHandler[Model any](ctx *gin.Context) {
	var o Model
	if err := ctx.ShouldBindBodyWithJSON(&o); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    err.Error(),
			Data:   nil,
		})
		return
	}
	resp := service.Insert[Model](o)
	ctx.JSON(http.StatusOK, resp)
	return
}
