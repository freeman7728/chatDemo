package idl

import (
	"github.com/gin-gonic/gin"
)

type IUser interface {
	UserRegister(ctx *gin.Context)
}
