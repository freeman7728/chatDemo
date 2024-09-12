package api

import (
	_ "chat/idl"
	"chat/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserApi struct{}

func (UserApi) UserRegister(c *gin.Context) {
	//TODO 绑定参数
	var userRegisterService service.UserService

	if err := c.ShouldBindBodyWithJSON(&userRegisterService); err == nil {
		resp := userRegisterService.UserRegisterService()
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Error(err)
	}
	//TODO 传递对象给服务层
	return
}

func (UserApi) UserLogin(c *gin.Context) {
	//TODO 绑定参数
	var userLoginService service.UserService

	if err := c.ShouldBindBodyWithJSON(&userLoginService); err == nil {
		resp := userLoginService.Login()
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Error(err)
	}
	//TODO 传递对象给服务层
	return
}
